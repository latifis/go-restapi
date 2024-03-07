package models

import "time"

type User struct {
	Id        int64     `gorm:"primaryKey" json:"id"`
	Username  string    `gorm:"varchar(300)" json:"username"`
	Password  string    `gorm:"varchar(300)" json:"password"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
	DeletedAt time.Time `gorm:"default:null" json:"deleted_at"`
}
