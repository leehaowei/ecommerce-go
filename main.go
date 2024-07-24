package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/leehaowei/ecommerce-go/db"
	"github.com/leehaowei/ecommerce-go/handler"
	"github.com/leehaowei/ecommerce-go/middleware"
	"github.com/leehaowei/ecommerce-go/routes"
)

func main() {
	port := FetchEnv("PORT")
	if port == "" {
		port = "8000"
	}

	app := handler.NewApplication(
		db.ProductData(db.Client, "Products"),
		db.UserData(db.Client, "Users"),
		db.UserData(db.Client, "Order"))

	router := gin.New()
	router.Use(gin.Logger())

	routes.UserRoutes(router)
	router.Use(middleware.Authentication())

	router.GET("/addtocart", app.AddToCart())
	router.GET("/removeitem", app.RemoveItem())
	router.GET("/cartcheckout", app.BuyFromCart())
	router.GET("/indstantbut", app.InstantBuy())

	log.Fatal(router.Run(":" + port))
}

func FetchEnv(key string) string {
	value, exists := os.LookupEnv(key)

	if !exists {
		log.Fatalf("FATAL: Environment variable %s is not set!", key)
	}

	return value
}
