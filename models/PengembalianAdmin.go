package models

import (
	"github.com/jinzhu/gorm"
)

func PengembalianAdmin(db *gorm.DB, confirm *StatusRental, status *StatusRental, id string) error {
	updateFields := map[string]interface{}{
		"status": "Pengembalian Sukses!",
	}
	statusUpdate := map[string]interface{}{
		"status": "Available",
	}
	err := db.Model(&StatusRental{}).Where("id=?", id).Update(updateFields).Error
	if err != nil {
		return err
	}
	err = db.Select("pinjaman").First(status, id).Error
	if err != nil {
		return err
	}
	pinjaman := status.Pinjaman
	err = db.Model(&StatusListRental{}).Where("namamobil =?", pinjaman).Update(statusUpdate).Error
	if err != nil {

		return err
	}
	return nil
}
