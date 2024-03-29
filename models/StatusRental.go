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

func GetStatusRental(db *gorm.DB, status *StatusRental, ID string) error {
	return db.Where("id=?", ID).First(status).Error
}

func CreateStatusRental(db *gorm.DB, createStatus *StatusRental) error {
	return db.Create(createStatus).Error
}
