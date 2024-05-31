package models

import (
	"time"
)

type Item struct {
	Kode      string    `json:"kode" gorm:"primaryKey; type:varchar(10)"`
	Nama      string    `json:"nama" gorm:"type:varchar(20); not null"`
	Jumlah    uint      `json:"jumlah" gorm:"type: int"`
	Deskripsi string    `json:"deskripsi" gorm:"type:text"`
	Status    bool      `json:"status" gorm:"default:true"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
