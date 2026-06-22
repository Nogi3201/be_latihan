package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	err := godotenv.Load(".env")
	if err != nil {
		err = godotenv.Load("../.env")
		if err != nil {
			log.Println(".env tidak ditemukan")
		}
	}

	dsn := os.Getenv("SUPABASE_DSN")
	if dsn == "" {
		log.Fatal("SUPABASE_DSN tidak ditemukan didalam file .env")
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Gagal terhubung ke database: %v", err)
	}

	DB = db
	log.Println("Koneksi ke PostgreSQL berhasil")
}

func GetDB() *gorm.DB {
	if DB == nil {
		log.Fatal("DB belum di inisialisasi")
	}
	return DB
}
