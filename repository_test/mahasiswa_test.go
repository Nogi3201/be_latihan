package repository_test

import (
	"be_latihan/config"
	"be_latihan/model"
	"be_latihan/repository"
	"fmt"
	"testing"
	"time"
)

func setupTest(t *testing.T) {
	config.InitDB()

	err := config.GetDB().AutoMigrate(&model.Mahasiswa{})
	if err != nil {
		t.Fatalf("Gagal migrasi database: %v", err)
	}
}

func TestInsertMahasiswa(t *testing.T) {
	setupTest(t)

	npm := time.Now().UnixNano()

	mhs := model.Mahasiswa{
		NPM:    npm,
		Nama:   "Malik",
		Prodi:  "Teknik Informatika",
		Alamat: "Bandung",
		Email:  "TestUpdate@gmail.com",
		NoHP:   "083813456789",
		Hobi:   []string{"Bermain Drum", "Game"},
	}

	_, err := repository.InsertMahasiswa(&mhs)
	if err != nil {
		t.Errorf("Insert failed: %v", err)
	}
	fmt.Printf("INSERTED NPM: %d\n", npm)
}

func TestGetAllMahasiswa(t *testing.T) {
	setupTest(t)

	data, err := repository.GetAllMahasiswa()
	if err != nil {
		t.Errorf("GetAll Gagal : %v", err)
	}

	if len(data) == 0 {
		t.Errorf("Data tidak ditemukan")
	}

	fmt.Printf("Berhasil. Data di table: %v\n", data)
}

func TestGetMahasiswaByNPM(t *testing.T) {
	setupTest(t)

	npm := int64(220001001)

	mhs, err := repository.GetMahasiswaByNPM(npm)
	if err != nil {
		t.Errorf("GetMahasiswaByNPM Gagal : %v", err)
	}

	if mhs.NPM != npm {
		t.Errorf("Expected NPM %d, got %d", npm, mhs.NPM)
	}

	fmt.Printf("Data Mahasiswa Ditemukan: %v\n", mhs)
}

func TestUpdateMahasiswa(t *testing.T) {
	setupTest(t)

	npm := int64(1775467813399742100)

	_, err := repository.UpdateMahasiswa(npm, &model.Mahasiswa{
		Nama:   "Muhammad Malik Nur",
		Prodi:  "D4 Teknik Informatika",
		Alamat: "Jl. Sukasari, Sukajadi, Kota Bandung",
		Email:  "Kamalputra1177@gmail.com",
		NoHP:   "082118937714",
		Hobi:   []string{"Bermain Alat Musik", "Bermain Game"},
	})

	if err != nil {
		t.Errorf("Update Failed : %v", err)
	}
	//fmt.Printf("Berhasil. NPM yang ditambahkan: %d\n", npm)

}

func TestDeleteMahasiswa(t *testing.T) {
	setupTest(t)

	npm := int64(1775467747236911100)

	err := repository.DeleteMahasiswa(npm)
	if err != nil {
		t.Errorf("Delete Failed : %v", err)
	}

}
