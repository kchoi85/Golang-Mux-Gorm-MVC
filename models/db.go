package model

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"kchoi85.io/rest-api/configs"
	"log"
)

var DB *gorm.DB

func InitDB() {
	config, err := util.LoadConfig("./configs/")
	if err != nil {
		log.Fatal("cannot load config: ", err)
	}

	DB, err = gorm.Open(postgres.Open(config.DBSource), &gorm.Config{})
	if err != nil {
		log.Fatal("cannot connect to db: ", err)
	}

	DB.AutoMigrate(&Book{})
	DB.AutoMigrate(&Author{})
}
