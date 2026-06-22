package model

type User struct {
	ID       string `json:"id" gorm:"column:id;type:uuid;default:gen_random_uuid();primaryKey" example:"123e4567-e89b-12d3-a456-426614174000"`
	Username string `json:"username" gorm:"column:username;type:varchar(50);uniqueIndex;not null" example:"admin"`
	Password string `json:"-" gorm:"column:password;type:varchar(255);not null"`
	Role     string `json:"role" gorm:"column:role;type:varchar(20);not null;default:admin" example:"admin"`
}

func (User) TableName() string { return "users" }

type AuthRequest struct {
	Username string `json:"username" example:"admin"`
	Password string `json:"password" example:"password123"`
	Role     string `json:"role" example:"admin"`
}

type AuthUserResponse struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Role     string `json:"role"`
}

type LoginResponse struct {
	Token string           `json:"token" example:"eyJhbGciOiJIUzI1NiIsInR5cCI..."`
	User  AuthUserResponse `json:"user"`
}

type ChangePasswordRequest struct {
	PasswordLama string `json:"password_lama"`
	PasswordBaru string `json:"password_baru"`
}
