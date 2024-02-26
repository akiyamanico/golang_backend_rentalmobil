package models

import (
	"github.com/jinzhu/gorm"
)

func Pengembalian(db *gorm.DB, confirm *StatusRental, id string) error {
	updateFields := map[string]interface{}{
		"status": "Pengembalian Menunggu Konfirmasi",
	}
	err := db.Model(&StatusRental{}).Where("id=?", id).Update(updateFields).Error
	if err != nil {
		return err
	}
	return nil
}
