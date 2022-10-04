package routers

import (
	"github.com/vikaputri/Gin_Golang/controllers"

	"github.com/gin-gonic/gin"
)

func StartServer() *gin.Engine {
	router := gin.Default()
	router.POST("/cars", controllers.CreateCar)
	router.POST("/cars/:carID", controllers.UpdateCar)
	router.POST("/cars/:carID", controllers.GetCar)
	router.POST("/cars/:carID", controllers.DeleteCar)
	return router
}
