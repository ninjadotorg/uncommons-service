package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/ninjadotorg/uncommons-service/models"
)

// DappController ...
type DappController struct{}

// Create ...
func (a DappController) Create(c *gin.Context) {
	dapp := models.Dapp{DappID: "aa"}
	db := models.Database()
	errDb := db.Create(&dapp).Error

}
