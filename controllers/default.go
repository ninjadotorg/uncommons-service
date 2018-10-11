package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// DefaultController : struct
type DefaultController struct{}

// Home ...
func (d DefaultController) Home(c *gin.Context) {
	resp := JSONResponse{1, "Uncommon REST API", c.Request.RemoteAddr}
	c.JSON(http.StatusOK, resp)
}

// NotFound ...
func (d DefaultController) NotFound(c *gin.Context) {
	resp := JSONResponse{0, "Page not found", nil}
	c.JSON(http.StatusOK, resp)
}
