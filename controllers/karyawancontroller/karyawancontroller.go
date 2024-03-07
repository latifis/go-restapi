package karyawancontroller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"go-restapi/helper"
	"go-restapi/models"

	"github.com/gorilla/mux"
)

// GetKaryawan mengambil daftar semua karyawan
func GetKaryawan(w http.ResponseWriter, r *http.Request) {
	var karyawan []models.Karyawan
	result := models.DB.Find(&karyawan)
	if result.Error != nil {
		helper.ResponseError(w, http.StatusInternalServerError, "Failed to fetch karyawan data")
		return
	}

	// Filter data yang telah dihapus secara lunak (soft deleted)
	var filteredKaryawan []models.Karyawan
	for _, k := range karyawan {
		if k.DeletedAt.IsZero() {
			filteredKaryawan = append(filteredKaryawan, k)
		}
	}

	if len(filteredKaryawan) == 0 {
		helper.ResponseError(w, http.StatusNotFound, "No active karyawan found")
		return
	}

	helper.ResponseJSON(w, http.StatusOK, filteredKaryawan)
}

// GetKaryawanByID mengambil karyawan berdasarkan ID
func GetKaryawanByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	karyawanID, err := strconv.Atoi(params["id"])
	if err != nil {
		helper.ResponseError(w, http.StatusBadRequest, "Invalid karyawan ID")
		return
	}

	var karyawan models.Karyawan
	result := models.DB.First(&karyawan, karyawanID)
	if result.Error != nil {
		helper.ResponseError(w, http.StatusNotFound, "Karyawan not found")
		return
	}

	// Periksa apakah karyawan telah dihapus secara lunak (soft deleted)
	if !karyawan.DeletedAt.IsZero() {
		helper.ResponseError(w, http.StatusNotFound, "Karyawan not found")
		return
	}

	helper.ResponseJSON(w, http.StatusOK, karyawan)
}

// CreateKaryawan membuat karyawan baru
func CreateKaryawan(w http.ResponseWriter, r *http.Request) {
	var karyawan models.Karyawan
	err := json.NewDecoder(r.Body).Decode(&karyawan)
	if err != nil {
		fmt.Println(err)
		helper.ResponseError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	models.DB.Create(&karyawan)
	helper.ResponseJSON(w, http.StatusCreated, karyawan)
}

// UpdateKaryawan mengubah data karyawan berdasarkan ID
func UpdateKaryawan(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	karyawanID, err := strconv.Atoi(params["id"])
	if err != nil {
		helper.ResponseError(w, http.StatusBadRequest, "Invalid karyawan ID")
		return
	}

	var karyawan models.Karyawan
	result := models.DB.First(&karyawan, karyawanID)
	if result.Error != nil {
		helper.ResponseError(w, http.StatusNotFound, "Karyawan not found")
		return
	}

	// Cek apakah karyawan telah dihapus secara lunak (soft deleted)
	if !karyawan.DeletedAt.IsZero() {
		helper.ResponseError(w, http.StatusNotFound, "Karyawan has been soft deleted")
		return
	}

	var updatedKaryawan models.Karyawan
	err = json.NewDecoder(r.Body).Decode(&updatedKaryawan)
	if err != nil {
		helper.ResponseError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	// Atur hanya bidang yang ingin diubah
	karyawan.Nama = updatedKaryawan.Nama
	karyawan.Nip = updatedKaryawan.Nip
	karyawan.TempatLahir = updatedKaryawan.TempatLahir
	karyawan.TanggalLahir = updatedKaryawan.TanggalLahir
	karyawan.Alamat = updatedKaryawan.Alamat
	karyawan.Agama = updatedKaryawan.Agama
	karyawan.JenisKelamin = updatedKaryawan.JenisKelamin
	karyawan.NoHandphone = updatedKaryawan.NoHandphone
	karyawan.Email = updatedKaryawan.Email

	models.DB.Save(&karyawan)
	helper.ResponseJSON(w, http.StatusOK, karyawan)
}

// DeleteKaryawan menghapus karyawan berdasarkan ID
func DeleteKaryawan(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	karyawanID, err := strconv.Atoi(params["id"])
	if err != nil {
		helper.ResponseError(w, http.StatusBadRequest, "Invalid karyawan ID")
		return
	}

	var karyawan models.Karyawan
	result := models.DB.First(&karyawan, karyawanID)
	if result.Error != nil {
		helper.ResponseError(w, http.StatusNotFound, "Karyawan not found")
		return
	}

	// Melakukan soft delete dengan mengatur deleted_at menjadi waktu saat ini
	models.DB.Model(&karyawan).Update("deleted_at", time.Now())
	helper.ResponseJSON(w, http.StatusOK, map[string]string{"message": "Karyawan soft deleted successfully"})
}
