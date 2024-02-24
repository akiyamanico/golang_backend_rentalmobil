package models

import (
	"github.com/jinzhu/gorm"
)

type StatusRental struct {
	ID              uint   `json:"id" gorm:"primary_key"`
	Username        string `json:"username"`
	Pinjaman        string `json:"pinjaman"`
	Status          string `json:"status"`
	Dokumenpinjaman string `json:"dokumenpinjaman"`
}

func GetStatusRental(db *gorm.DB, status *[]StatusRental) error {
	return db.Find(status).Error
}

func UpdateStatusRental(db *gorm.DB, statusRentalID string, statusRental *StatusRental) error {
	return db.Model(&StatusRental{}).Where("id =?", statusRentalID).Update(statusRental).Error
}

func CreateStatusRental(db *gorm.DB, createStatus *StatusRental) error {
	return db.Create(createStatus).Error
}
