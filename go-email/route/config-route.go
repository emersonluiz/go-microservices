package route

import (
	"github.com/emersonluiz/go-email/controller"
	"github.com/gin-gonic/gin"
)

func ConfigRoute(router *gin.Engine) *gin.Engine {
	main := router.Group("email/v1")
	{
		check := main.Group("checks")
		{
			check.GET("/", controller.Check)
		}
	}

	return router
}
