package routes

import(
	"github.com/gin-gonic/gin"
	controller "github.com/andymartinezot/go-restaurant-management/server/internal/controllers"
)

func FoodRoutes(incomingRoutes *gin.Engine){
	incomingRoutes.GET("/foods", controller.GetFoods())
	incomingRoutes.GET("/foods/:food_id", controller.GetFood())
	incomingRoutes.POST("foods", controller.CreateFood())
	incomingRoutes.PATCH("foods/:foods_id", controller.UpdateFood())
}