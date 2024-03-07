package models

import "time"

type Karyawan struct {
	Id           int64     `gorm:"primaryKey" json:"id"`
	Nama         string    `gorm:"varchar(300)" json:"nama"`
	Nip          string    `gorm:"varchar(300)" json:"nip"`
	TempatLahir  string    `gorm:"varchar(300)" json:"tempat_lahir"`
	TanggalLahir time.Time `gorm:"type:date" json:"tanggal_lahir"`
	Umur         int64     `gorm:"-" json:"umur"`
	Alamat       string    `gorm:"varchar(300)" json:"alamat"`
	Agama        string    `gorm:"varchar(300)" json:"agama"`
	JenisKelamin string    `gorm:"varchar(300)" json:"jenis_kelamin"`
	NoHandphone  string    `gorm:"varchar(300)" json:"no_handphone"`
	Email        string    `gorm:"varchar(300)" json:"email"`
	CreatedAt    time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt    time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
	DeletedAt    time.Time `gorm:"default:null" json:"deleted_at"`
}
