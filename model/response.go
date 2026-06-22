package model

type Response struct {
	Message string      `json:"message" example:"pesan balasan"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty" example:"pesan error jika ada"`
}

type UnauthorizedResponse struct {
	Message string `json:"message" example:"token tidak valid"`
}

type ForbiddenResponse struct {
	Message string `json:"message" example:"user tidak memiliki akses untuk fitur ini"`
}

type SuccessResponse struct {
	Message string `json:"message" example:"berhasil"`
}

type CreatedResponse struct {
	Message string `json:"message" example:"data berhasil dibuat"`
}
