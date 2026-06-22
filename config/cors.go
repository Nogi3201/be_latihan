package config

var allowedOrigins = []string{
	"http://localhost:5173",
	"https://belatihan-production-a189.up.railway.app",
	"my-iqilpt53z-study3201.vercel.app",
}

func GetAllowedOrigins() []string {
	return allowedOrigins
}