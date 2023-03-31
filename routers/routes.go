package routers

import (
	"Gawean/controllers"
	"Gawean/middleware"

	"github.com/gin-gonic/gin"
)

func StartServer() *gin.Engine {
	router := gin.Default()
	clients := router.Group("/clients")
	clients.GET("", middleware.BasicAuth, controllers.GetAllClients)
	clients.POST("", middleware.BasicAuth, controllers.InsertClient)
	clients.PUT("/:id", middleware.BasicAuth, controllers.UpdateClient)
	clients.DELETE("/:id", middleware.BasicAuth, controllers.DeleteClient)
	category := router.Group("/category")
	category.GET("", middleware.BasicAuth, controllers.GetAllCategories)
	category.POST("", middleware.BasicAuth, controllers.InsertCategory)
	category.PUT("/:id", middleware.BasicAuth, controllers.UpdateCategory)
	category.DELETE("/:id", middleware.BasicAuth, controllers.DeleteCategory)
	address := router.Group("/address")
	address.GET("", middleware.BasicAuth, controllers.GetAllAddresses)
	address.POST("", middleware.BasicAuth, controllers.InsertAddress)
	address.PUT("/:id", middleware.BasicAuth, controllers.UpdateAddress)
	address.DELETE("/:id", middleware.BasicAuth, controllers.DeleteAddress)
	gawean := router.Group("/gawean")
	gawean.GET("", middleware.BasicAuth, controllers.GetAllGawean)
	gawean.POST("", middleware.BasicAuth, controllers.InsertGawean)
	gawean.PUT("/:id", middleware.BasicAuth, controllers.UpdateGawean)
	gawean.DELETE("/:id", middleware.BasicAuth, controllers.DeleteGawean)
	return router
}