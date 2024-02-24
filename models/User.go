package models

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	ID       uint   `json:"id" gorm:"primary_key"`
	Username string `json:"username"`
	Password string `json:"password"`
	Level    string `json:"level"`
}

func GetAllUsers(db *gorm.DB, users *[]User) error {
	return db.Find(users).Error
}

func GetUserByID(db *gorm.DB, user *User, userID string) error {
	return db.Where("id=?", userID).First(user).Error
}

func CreateUser(db *gorm.DB, user *User) error {
	return db.Create(user).Error
}
