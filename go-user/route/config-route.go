package route

import (
	"github.com/emersonluiz/go-user/controller"
	"github.com/gin-gonic/gin"
)

func ConfigRoute(router *gin.Engine) *gin.Engine {
	main := router.Group("user/v1")
	{
		check := main.Group("checks")
		{
			check.GET("/", controller.Check)
		}
		user := main.Group("users")
		{
			user.POST("/", controller.CreateUser)
			user.GET("/", controller.FindAllUser)
			user.DELETE("/:id", controller.DeleteUser)
			user.GET("/:id", controller.FindOneUser)
			user.PUT("/:id", controller.CreateUser)
		}
	}

	return router
}
