package main

// @title API Praktikum 13
// @version 1.0
// @description Dokumentasi API Golang Fiber
// @host 127.0.0.1:3000
// @BasePath /
// @schemes http https
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization

import (
	"be_latihan/config"
	"be_latihan/docs"
	"be_latihan/model"
	"be_latihan/router"
	"os"
	"strings"

	_ "be_latihan/docs"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	app := fiber.New()
	app.Use(logger.New())

	// ✅ CORS HARUS DI ATAS
	app.Use(cors.New(cors.Config{
		AllowOrigins: strings.Join(config.GetAllowedOrigins(), ","),
		AllowMethods: "GET,POST,PUT,DELETE,OPTIONS",
		AllowHeaders: "Origin,Content-Type,Accept,Authorization",
	}))

	swaggerHost := os.Getenv("SWAGGER_HOST")
	if swaggerHost == "" {
		swaggerHost = "localhost:3000"
	}
	docs.SwaggerInfo.Host = swaggerHost

	config.InitDB()

	// automigrate
	config.GetDB().AutoMigrate(&model.Mahasiswa{}, &model.User{})

	router.SetupRoutes(app)

	app.Listen(":3000")
}
