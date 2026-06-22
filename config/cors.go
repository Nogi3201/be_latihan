package config

var allowedOrigins = []string{
	"http://localhost:5173",
	"belatihan-production-a189.up.railway.app",
}

func GetAllowedOrigins() []string {
	return allowedOrigins
}
