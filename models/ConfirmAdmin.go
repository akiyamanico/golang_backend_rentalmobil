package models

import "github.com/jinzhu/gorm"

type Confirm struct {
	ID              uint   `json:"id" gorm:"primary_key"`
	Username        string `json:"username"`
	Pinjaman        string `json:"pinjaman"`
	Status          string `json:"status"`
	Dokumenpinjaman string `json:"dokumenpinjaman"`
}

func UpdateConfirm(db *gorm.DB, confirm *Confirm, status *StatusRental, id string) error {
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
