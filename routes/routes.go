package routes

import (
	controllers "mrp-service/controllers"

	"log"

	"github.com/gin-gonic/gin"
)

func MainRoutes(r *gin.Engine) {
	r.POST("/login", controllers.Login)
	r.POST("/todo", controllers.CreateTodo)
	r.POST("/logout", controllers.Logout)
	r.POST("/refresh", controllers.Refresh)
	log.Fatal(r.Run(":8080"))
}
