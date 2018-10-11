package controllers

import (
	"encoding/hex"
	"net/http"

	"github.com/ninjadotorg/uncommons-service/models"
	"github.com/ninjadotorg/uncommons-service/services"
	"github.com/ninjadotorg/uncommons-service/utils"

	"github.com/gin-gonic/gin"
)

// UserController : struct
type UserController struct{}

// SignUp ...
func (u UserController) SignUp(c *gin.Context) {
	uncommonsWalletAddress := c.DefaultPostForm("uncommons_address", "")
	uncommonsSignedHash := c.DefaultPostForm("uncommons_signed_hash", "")
	uncommonsText := c.DefaultPostForm("uncommons_text", "")

	if len(uncommonsWalletAddress) == 0 || len(uncommonsSignedHash) == 0 || len(uncommonsText) == 0 {
		resp := JSONResponse{0, "Missing parameters", nil}
		c.JSON(http.StatusOK, resp)
		return
	}

	str, _ := hex.DecodeString(uncommonsText)
	res := utils.VerifySig(uncommonsWalletAddress, uncommonsSignedHash, []byte(str))
	if !res {
		resp := JSONResponse{0, "Cannot verify this wallet address", nil}
		c.JSON(http.StatusOK, resp)
		return
	}

	dispatcher := services.DispatcherService{}
	passpharse, result := dispatcher.SignUp()
	if !result {
		resp := JSONResponse{0, "Cannot signup", nil}
		c.JSON(http.StatusOK, resp)
		return
	}

	db := models.Database()
	user := models.User{Payload: uncommonsWalletAddress, WalletAddress: passpharse}
	errDb := db.Create(&user).Error
	if errDb != nil {
		resp := JSONResponse{0, "Sign up failed", nil}
		c.JSON(http.StatusOK, resp)
		return
	}

	resp := JSONResponse{1, "", map[string]interface{}{"passpharse": passpharse}}
	c.JSON(http.StatusOK, resp)
	return
}
