package routers

import (
	"assignment2.id/orderapi/controllers"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func StartServer() *gin.Engine {
	router := gin.Default()
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.GET("/orders/:orderID", controllers.GetOrder)
	router.PUT("/orders/:orderID", controllers.UpdateOrder)
	router.POST("/orders", controllers.CreateOrder)
	router.DELETE("/orders/:orderID", controllers.DeleteOrder)
	return router
}
