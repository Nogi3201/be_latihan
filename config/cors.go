package config

var allowedOrigins = []string{
	"http://localhost:5173",
	"https://belatihan-production-a189.up.railway.app",
	"https://my-fe-tawny.vercel.app",
}

func GetAllowedOrigins() []string {
	return allowedOrigins
}
