package config

var allowedOrigins = []string{
	"http://localhost:5173",
	"https://belatihan-production-a189.up.railway.app",
	"https://belatihan-production-a189.up.railway.app",
}

func GetAllowedOrigins() []string {
	return allowedOrigins
}
