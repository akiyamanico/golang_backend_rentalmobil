package models

import "github.com/jinzhu/gorm"

type CreateRental struct {
	ID              uint   `json:"id" gorm:"primary_key" form:"id"`
	Username        string `json:"username"`
	Pinjaman        string `json:"pinjaman"`
	Status          string `json:"status"`
	Dokumenpinjaman string `json:"dokumenpinjaman"`
}

func CreateRentalMobil(db *gorm.DB, create *CreateRental) error {
	return db.Create(create).Error
}
