package handler

import (
	"be_latihan/config/middleware"
	"be_latihan/model"
	"be_latihan/pkg/password"
	"be_latihan/repository"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// Register godoc
// @Summary Register a new user
// @Description Register a new user with username and password
// @Tags Auth
// @Accept json
// @Produce json
// @Param request body model.AuthRequest true "Register request"
// @Success 201 {object} model.CreatedResponse
// @Failure 400 {object} model.Response
// @Failure 409 {object} model.Response
// @Failure 500 {object} model.Response
// @Router /register [post]
func Register(c *fiber.Ctx) error {
	var payload model.AuthRequest
	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(model.Response{
			Message: "payload tidak valid",
			Error:   err.Error(),
		})
	}

	payload.Username = strings.TrimSpace(payload.Username)
	payload.Role = strings.TrimSpace(payload.Role)
	if payload.Role == "" {
		payload.Role = "admin"
	}

	if payload.Username == "" || payload.Password == "" {
		return c.Status(fiber.StatusBadRequest).JSON(model.Response{
			Message: "username dan password wajib diisi",
		})
	}

	hashedPassword, err := password.HashPassword(payload.Password)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(model.Response{
			Message: "gagal membuat hash password",
			Error:   err.Error(),
		})
	}

	user := model.User{
		Username: payload.Username,
		Password: hashedPassword,
		Role:     payload.Role,
	}

	data, err := repository.InsertUser(&user)
	if err != nil {
		return c.Status(fiber.StatusConflict).JSON(model.Response{
			Message: "username sudah digunakan atau data tidak valid",
			Error:   err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(model.Response{
		Message: "register berhasil",
		Data: model.AuthUserResponse{
			ID:       data.ID,
			Username: data.Username,
			Role:     data.Role,
		},
	})
}

// Login godoc
// @Summary Login user
// @Description Login with username and password to get JWT token
// @Tags Auth
// @Accept json
// @Produce json
// @Param request body model.AuthRequest true "Login request"
// @Success 200 {object} model.LoginResponse
// @Failure 400 {object} model.Response
// @Failure 401 {object} model.UnauthorizedResponse
// @Failure 500 {object} model.Response
// @Router /login [post]
func Login(c *fiber.Ctx) error {
	var payload model.AuthRequest
	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(model.Response{
			Message: "payload tidak valid",
			Error:   err.Error(),
		})
	}

	user, err := repository.FindUserByUsername(strings.TrimSpace(payload.Username))
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusUnauthorized).JSON(model.Response{
				Message: "username atau password salah",
			})
		}

		return c.Status(fiber.StatusInternalServerError).JSON(model.Response{
			Message: "gagal mencari user",
			Error:   err.Error(),
		})
	}

	if !password.CheckPasswordHash(payload.Password, user.Password) {
		return c.Status(fiber.StatusUnauthorized).JSON(model.Response{
			Message: "username atau password salah",
		})
	}

	token, err := middleware.GenerateJWT(user, 2*time.Hour)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(model.Response{
			Message: "gagal membuat token",
			Error:   err.Error(),
		})
	}

	return c.JSON(model.Response{
		Message: "login berhasil",
		Data: model.LoginResponse{
			Token: token,
			User: model.AuthUserResponse{
				ID:       user.ID,
				Username: user.Username,
				Role:     user.Role,
			},
		},
	})
}

// ChangePassword godoc
// @Summary Change user password
// @Description Change the password of the currently authenticated user
// @Tags Auth
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body model.ChangePasswordRequest true "Change Password Request"
// @Success 200 {object} model.SuccessResponse
// @Failure 400 {object} model.Response
// @Failure 401 {object} model.UnauthorizedResponse
// @Failure 404 {object} model.Response
// @Failure 500 {object} model.Response
// @Router /api/change-password [post]
func ChangePassword(c *fiber.Ctx) error {
	username := c.Locals("username").(string)

	var payload model.ChangePasswordRequest
	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(model.Response{
			Message: "payload tidak valid",
			Error:   err.Error(),
		})
	}

	user, err := repository.FindUserByUsername(username)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(model.Response{
			Message: "user tidak ditemukan",
		})
	}

	if !password.CheckPasswordHash(payload.PasswordLama, user.Password) {
		return c.Status(fiber.StatusUnauthorized).JSON(model.Response{
			Message: "password lama salah",
		})
	}

	hashedPassword, err := password.HashPassword(payload.PasswordBaru)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(model.Response{
			Message: "gagal membuat hash password baru",
		})
	}

	err = repository.UpdateUserPassword(username, hashedPassword)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(model.Response{
			Message: "gagal update password",
			Error:   err.Error(),
		})
	}

	return c.JSON(model.Response{
		Message: "password berhasil diubah",
	})
}
