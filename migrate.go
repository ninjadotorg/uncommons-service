package main

import (
	"github.com/jinzhu/gorm"
	"github.com/ninjadotorg/uncommons-service/config"
	"github.com/ninjadotorg/uncommons-service/models"
)

func main() {
	config.Init()

	//
	var db *gorm.DB = models.Database()
	defer db.Close()

	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Dapp{})
}
