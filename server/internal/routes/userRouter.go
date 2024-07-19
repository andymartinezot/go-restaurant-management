package routes

import(
	"github.com/gin-gonic/gin"
	controller "github.com/andymartinezot/go-restaurant-management/server/internal/controllers"
)

func UserRoutes(incomingRoutes *gin.Engine){
	incomingRoutes.GET("/users", controller.GetUsers())
	incomingRoutes.GET("/users/:user_id", controller.GetUser())
	incomingRoutes.POST("/users/signup", controller.SignUp())
	incomingRoutes.PORT("/users/login", controller.Login())
}