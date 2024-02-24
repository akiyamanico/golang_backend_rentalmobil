package config

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Config struct {
	Host     string
	Port     string
	User     string
	DBName   string
	Password string
	DB       *gorm.DB
}

func Build() *Config {
	config := Config{
		Host:     "localhost",
		Port:     "3306",
		User:     "root",
		DBName:   "rentalmobil",
		Password: "",
	}

	db, err := gorm.Open("mysql", config.URL())
	if err != nil {
		panic("Failed to connect to the database")
	}
	db.LogMode(true)
	config.DB = db
	return &config
}

func (c *Config) URL() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		c.User,
		c.Password,
		c.Host,
		c.Port,
		c.DBName)
}
