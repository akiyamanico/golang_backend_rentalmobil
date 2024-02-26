package models

import "github.com/jinzhu/gorm"

func CreateRentalMobil(db *gorm.DB, create *StatusRental) error {
	return db.Create(create).Error
}
