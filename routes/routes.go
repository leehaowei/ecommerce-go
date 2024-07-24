package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/leehaowei/ecommerce-go/handler"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("/users/signup", handler.SignUp())
	incomingRoutes.POST("/users/login", handler.Login())
	incomingRoutes.POST("/admin/addproduct", handler.ProductViewerAdmin())
	incomingRoutes.GET("/users/productview", handler.SearchProduct())
	incomingRoutes.GET("/users/search", handler.SearchProductByQuery())
}