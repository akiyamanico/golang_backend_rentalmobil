package models

import "github.com/jinzhu/gorm"

func UpdateConfirm(db *gorm.DB, confirm *StatusRental, status *StatusRental, id string) error {
	updateFields := map[string]interface{}{
		"status": "Sudah Dikonfirmasi!",
	}
	err := db.Model(&StatusRental{}).Where("id =?", id).Update(updateFields).Error
	if err != nil {
		return err
	}

	carUpdate := map[string]interface{}{
		"status": "In Use",
	}
	err = db.Select("pinjaman").First(status, id).Error
	if err != nil {
		return err
	}
	pinjaman := status.Pinjaman
	err = db.Model(&StatusListRental{}).Where("namamobil =?", pinjaman).Update(carUpdate).Error
	if err != nil {
		return err
	}

	return nil
}
