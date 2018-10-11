package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// UserController : struct
type UserController struct{}

// SignUp ...
func (u UserController) SignUp(c *gin.Context) {
	uncommonsWalletAddres := c.DefaultPostForm("uncommons_address", "")
	uncommonsSignedHash := c.DefaultPostForm("uncommons_signed_hash", "")
	uncommonsText := c.DefaultPostForm("uncommons_text", "")

	if len(uncommonsWalletAddres) == 0 || len(uncommonsSignedHash) == 0 || len(uncommonsText) == 0 {
		resp := JSONResponse{0, "Missing parameters", nil}
		c.JSON(http.StatusOK, resp)
		return
	}

	resp := JSONResponse{1, "", map[string]interface{}{"passpharse": ""}}
	c.JSON(http.StatusOK, resp)
	return
}
