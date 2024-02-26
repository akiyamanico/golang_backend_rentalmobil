package models

import "github.com/jinzhu/gorm"

func GetMobil(db *gorm.DB, mobil *[]StatusListRental) error {
	return db.Where("status = ?", "Available").Find(&mobil).Error
}
