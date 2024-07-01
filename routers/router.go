package routers

import (
	setting "ginblog/utils"
	"github.com/gin-gonic/gin"
)

func InitRouter() {
	gin.SetMode(setting.AppMode)
	r := gin.Default()

	router := r.Group("api/v1")
	{
		router.GET("hello", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "hello world",
			})
		})
	}

	err := r.Run(setting.HttpPort)
	if err != nil {
		return
	}

}
