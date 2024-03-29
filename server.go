package server

import (
	"AustemServer/handlers/handlers"
	"github.com/gin-gonic/gin"
)

func StartAPI() {

	r := gin.Default()

	productGroup := r.Group("/products")
	{
		productGroup.GET("/getall", handlers.GetAllProducts)
		productGroup.POST("/add", handlers.PostProduct)

	}
	userGroup := r.Group("/user")
	{
		userGroup.POST("/auth", handlers.Autharization)
	}
	projectGroup := r.Group("/project")
	{
		projectGroup.GET("/getall", handlers.Autharization)
		projectGroup.GET("/get", handlers.Autharization)
		projectGroup.POST("/add", handlers.Autharization)
	}
	err := r.Run(":8080")
	if err != nil {
		return
	}
}
