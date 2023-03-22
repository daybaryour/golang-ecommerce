package routes

import (
	"github.com/daybaryour/golang-ecommerce/pkg/controllers"
	"github.com/gin-gonic/gin"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("/users/signup", controllers.Signup())
	incomingRoutes.POST("/users/login", controllers.Login())
	incomingRoutes.POST("/admin/add_product", controllers.AddProduct())
	incomingRoutes.GET("/users/products", controllers.GetProducts())
	incomingRoutes.GET("/users/search", controllers.searchProduct())
}
