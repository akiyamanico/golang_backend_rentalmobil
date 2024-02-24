package models

import (
	"strconv"

	"github.com/jinzhu/gorm"
)

type StatusListRental struct {
	ID        uint   `json:"id" gorm:"primary_key"`
	Namamobil string `json:"namamobil"`
	Keadaan   string `json:"keadaan"`
	Status    string `json:"status"`
}

func GetStatusListRental(db *gorm.DB, page string, pageSize string) ([]StatusListRental, error) {
	var status []StatusListRental
	pageInt, _ := strconv.Atoi(page)
	pageSizeInt, _ := strconv.Atoi(pageSize)
	offset := (pageInt - 1) * pageSizeInt

	err := db.Limit(pageSizeInt).Offset(offset).Find(&status).Error
	if err != nil {
		return nil, err
	}
	return status, nil
}
