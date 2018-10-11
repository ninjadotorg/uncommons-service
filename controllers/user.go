package controllers

import (
	"net/http"

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

	res := utils.VerifySig(uncommonsWalletAddress, uncommonsSignedHash, []byte(uncommonsText))
	if !res {
		resp := JSONResponse{0, "Invalid parameters", nil}
		c.JSON(http.StatusOK, resp)
		return
	}

	resp := JSONResponse{1, "", map[string]interface{}{"passpharse": ""}}
	c.JSON(http.StatusOK, resp)
	return
}
