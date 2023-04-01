package routers

import (
	"Gawean/controllers"
	"Gawean/middleware"

	"github.com/gin-gonic/gin"
)

func StartServer() *gin.Engine {
	router := gin.Default()
	router.POST("/register", controllers.RegisterUser)

	clients := router.Group("/clients")
	clients.GET("", middleware.BasicAuth, controllers.GetAllClients)
	clients.GET("/:id", middleware.BasicAuth, controllers.GetClient)
	clients.GET("/city/:id", middleware.BasicAuth, controllers.GetClientByCity)
	clients.POST("", middleware.BasicAuth, controllers.InsertClient)
	clients.PUT("/:id", middleware.BasicAuth, controllers.UpdateClient)
	clients.DELETE("/:id", middleware.BasicAuth, controllers.DeleteClient)
	clients.DELETE("/:id/:name", middleware.BasicAuth, controllers.DeleteClientSpecific)
	category := router.Group("/category")
	category.GET("", middleware.BasicAuth, controllers.GetAllCategories)
	category.POST("", middleware.BasicAuth, controllers.InsertCategory)
	category.PUT("/:id", middleware.BasicAuth, controllers.UpdateCategory)
	category.DELETE("/:id", middleware.BasicAuth, controllers.DeleteCategory)
	address := router.Group("/address")
	address.GET("", middleware.BasicAuth, controllers.GetAllAddresses)
	address.GET("/:id", middleware.BasicAuth, controllers.GetAddress)
	address.POST("", middleware.BasicAuth, controllers.InsertAddress)
	address.PUT("/:id", middleware.BasicAuth, controllers.UpdateAddress)
	address.DELETE("/:id", middleware.BasicAuth, controllers.DeleteAddress)
	gawean := router.Group("/gawean")
	gawean.GET("", middleware.BasicAuth, controllers.GetAllGawean)
	gawean.GET("/client/:id", middleware.BasicAuth, controllers.GetGaweanByClient)
	gawean.GET("/category/:id", middleware.BasicAuth, controllers.GetGaweanByCategory)
	gawean.POST("", middleware.BasicAuth, controllers.InsertGawean)
	gawean.PUT("/:id", middleware.BasicAuth, controllers.UpdateGawean)
	gawean.DELETE("/:id", middleware.BasicAuth, controllers.DeleteGawean)

	router.GET("/date", middleware.BasicAuth, controllers.GetByDate)
	return router
}
