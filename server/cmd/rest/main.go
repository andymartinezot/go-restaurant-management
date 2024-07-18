package main

import(
	"fmt"
	"os"
	"github.com/gin-gonic/gin"
	"github.com/andymartinezot/go-restaurant-management/server/internal/database"
	"github.com/andymartinezot/go-restaurant-management/server/internal/routes"
	"github.com/andymartinezot/go-restaurant-management/server/internal/middleware"
	"go.mongodb.org/mongo-driver/mongo"
)

var foodCollection *mongo.Collection = database.OpenCollection(database.Client, "food")

func main(){
	port := os.Getenv("PORT")

	if port == ""{
		port = "8080"
	}

	//Used for JWT authentication
	router := gin.New()
	router.Use(gin.Logger())
	routes.UserRoutes(router)
	router.Use(middleware.Authentication())

	routes.FoodRoutes(router)
	routes.InvoiceRoutes(router)
	routes.MenuRoutes(router)
	//routes.NoteRoutes(router)
	routes.OrderItemRoutes(router)
	routes.OrderRoutes(router)
	routes.TableRoutes(router)

	router.Run(":" + port)

}

