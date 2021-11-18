package controller

import "github.com/gin-gonic/gin"

func Check(c *gin.Context) {
	c.String(200, "ok")
}
