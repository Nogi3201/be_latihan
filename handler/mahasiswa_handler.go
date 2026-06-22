package handler

import (
	"be_latihan/model"
	"be_latihan/repository"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// GetAllMahasiswa godoc
// @Summary Get all mahasiswa
// @Description Retrieve a list of all mahasiswa
// @Tags Mahasiswa
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} model.Response
// @Failure 401 {object} model.UnauthorizedResponse
// @Failure 500 {object} model.Response
// @Router /api/mahasiswa [get]
func GetAllMahasiswa(c *fiber.Ctx) error {
	data, err := repository.GetAllMahasiswa()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(model.Response{
			Message: "Failed to retrieve data",
			Error:   err.Error(),
		})
	}

	return c.Status(200).JSON(model.Response{
		Message: "Data retrieved successfully",
		Data:    data,
	})
}

// GetMahasiswaByNPM godoc
// @Summary Get mahasiswa by NPM
// @Description Retrieve a single mahasiswa by its NPM
// @Tags Mahasiswa
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param npm path int true "Mahasiswa NPM"
// @Success 200 {object} model.Response
// @Failure 400 {object} model.Response
// @Failure 401 {object} model.UnauthorizedResponse
// @Failure 404 {object} model.Response
// @Failure 500 {object} model.Response
// @Router /api/mahasiswa/{npm} [get]
func GetMahasiswaByNPM(c *fiber.Ctx) error {
	//npmQuery := c.Params("npm")
	//if npmQuery == "" {
	//	return c.Status(fiber.StatusBadRequest).JSON(model.Response{
	//		Message: "NPM parameter tidak boleh kosong",
	//	})
	//}
	//npm, err := strconv.ParseInt(npmQuery, 10, 64)

	npm, err := strconv.ParseInt(c.Params("npm"), 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(model.Response{
			Message: "Invalid NPM format",
			Error:   err.Error(),
		})
	}

	mhs, err := repository.GetMahasiswaByNPM(npm)
	if err != nil {
		if err.Error() == "record not found" {
			return c.Status(fiber.StatusNotFound).JSON(model.Response{
				Message: "Mahasiswa not found",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(model.Response{
			Message: "Failed to retrieve data",
			Error:   err.Error(),
		})
	}

	return c.JSON(model.Response{
		Message: "Data retrieved successfully",
		Data:    mhs,
	})
}

// InsertMahasiswa godoc
// @Summary Insert new mahasiswa
// @Description Add a new mahasiswa to the database
// @Tags Mahasiswa
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body model.Mahasiswa true "Mahasiswa Data"
// @Success 201 {object} model.CreatedResponse
// @Failure 400 {object} model.Response
// @Failure 401 {object} model.UnauthorizedResponse
// @Failure 500 {object} model.Response
// @Router /api/mahasiswa [post]
func InsertMahasiswa(c *fiber.Ctx) error {
	var payload model.Mahasiswa
	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(model.Response{
			Message: "payload tidak valid",
			Error:   err.Error(),
		})
	}

	data, err := repository.InsertMahasiswa(&payload)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(model.Response{
			Message: "gagal menambahkan data mahasiswa",
			Error:   err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(model.Response{
		Message: "berhasil menambahkan data mahasiswa",
		Data:    data,
	})
}

// UpdateMahasiswa godoc
// @Summary Update mahasiswa data
// @Description Update existing mahasiswa data by NPM
// @Tags Mahasiswa
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param npm path int true "Mahasiswa NPM"
// @Param request body model.Mahasiswa true "Updated Mahasiswa Data"
// @Success 200 {object} model.SuccessResponse
// @Failure 400 {object} model.Response
// @Failure 401 {object} model.UnauthorizedResponse
// @Failure 404 {object} model.Response
// @Failure 500 {object} model.Response
// @Router /api/mahasiswa/{npm} [put]
func UpdateMahasiswa(c *fiber.Ctx) error {
	npm, err := strconv.ParseInt(c.Params("npm"), 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(model.Response{
			Message: "npm tidak valid",
			Error:   err.Error(),
		})
	}

	var payload model.Mahasiswa
	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(model.Response{
			Message: "payload tidak valid",
			Error:   err.Error(),
		})
	}

	data, err := repository.UpdateMahasiswa(npm, &payload)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(model.Response{
				Message: "data mahasiswa tidak ditemukan",
			})
		}

		return c.Status(fiber.StatusInternalServerError).JSON(model.Response{
			Message: "gagal mengubah data mahasiswa",
			Error:   err.Error(),
		})
	}

	return c.JSON(model.Response{
		Message: "berhasil mengubah data mahasiswa",
		Data:    data,
	})
}

// DeleteMahasiswa godoc
// @Summary Delete mahasiswa
// @Description Delete a mahasiswa from the database by NPM
// @Tags Mahasiswa
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param npm path int true "Mahasiswa NPM"
// @Success 200 {object} model.SuccessResponse
// @Failure 400 {object} model.Response
// @Failure 401 {object} model.UnauthorizedResponse
// @Failure 500 {object} model.Response
// @Router /api/mahasiswa/{npm} [delete]
func DeleteMahasiswa(c *fiber.Ctx) error {
	npm, err := strconv.ParseInt(c.Params("npm"), 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(model.Response{
			Message: "npm tidak valid",
			Error:   err.Error(),
		})
	}

	if err := repository.DeleteMahasiswa(npm); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(model.Response{
			Message: "gagal menghapus data mahasiswa",
			Error:   err.Error(),
		})
	}

	return c.JSON(model.Response{
		Message: "berhasil menghapus data mahasiswa",
	})
}
