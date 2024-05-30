package models

import "gorm.io/gorm"

type Item struct {
	gorm.Model
	Kode      string `json:"kode" gorm:"primaryKey; type:varchar(10)"`
	Nama      string `json:"nama" gorm:"type:varchar(20); not null"`
	Jumlah    uint   `json:"jumlah" gorm:"type: int"`
	Deskripsi string `json:"deskripsi" gorm:"type:text"`
	Status    bool   `gorm:"default:true"`
}
