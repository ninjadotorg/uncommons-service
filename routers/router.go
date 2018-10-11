package routers

import (
	"io"
	"log"
	"os"

	"github.com/gin-gonic/gin"

	"github.com/ninjadotorg/uncommons-service/controllers"
	"github.com/ninjadotorg/uncommons-service/middlewares"
)

func NewRouter() *gin.Engine {
	// Logger
	logFile, err := os.OpenFile("logs/log.txt", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		panic(err)
	}
	gin.DefaultWriter = io.MultiWriter(logFile, os.Stdout)
	log.SetOutput(gin.DefaultWriter) // You may need this
	log.SetFlags(log.Lshortfile | log.LstdFlags)

	router := gin.New()
	router.Use(gin.Logger())
	router.Use(middlewares.CORSMiddleware())
	router.Use(middlewares.ErrorHandler())

	defaultController := new(controllers.DefaultController)
	router.GET("/", defaultController.Home)

	userController := new(controllers.UserController)
	dappController := new(controllers.DappController)
	userGroup := router.Group("user")
	{
		userGroup.POST("/sign-up", userController.SignUp)
	}
	dappGroup := router.Group("dapp")
	{
		dappGroup.POST("/create", middlewares.AuthMiddleware(), dappController.Create)
	}

	router.NoRoute(defaultController.NotFound)
	return router
}
