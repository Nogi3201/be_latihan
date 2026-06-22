package model

import "github.com/lib/pq"

type Mahasiswa struct {
	NPM    int64          `json:"npm" gorm:"column:npm;primaryKey;type:bigint;not null" example:"123456789"`
	Nama   string         `json:"nama" gorm:"column:nama;type:varchar(100);not null" example:"Budi Santoso"`
	Prodi  string         `json:"prodi" gorm:"column:prodi;type:varchar(100);not null" example:"Teknik Informatika"`
	Alamat string         `json:"alamat" gorm:"column:alamat;type:varchar(200)" example:"Jl. Jendral Sudirman No. 1"`
	Email  string         `json:"email" gorm:"column:email;type:varchar(100)" example:"budi.santoso@example.com"`
	NoHP   string         `json:"no_hp" gorm:"column:no_hp;type:varchar(20)" example:"081234567890"`
	Hobi   pq.StringArray `json:"hobi" gorm:"column:hobi;type:text[]" swaggertype:"array,string" example:"membaca,menulis"`
}

func (Mahasiswa) TableName() string { return "mahasiswa" }
