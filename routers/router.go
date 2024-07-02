package routers

import (
	v1 "ginblog/api/v1"
	setting "ginblog/utils"
	"github.com/gin-gonic/gin"
)

func InitRouter() {
	gin.SetMode(setting.AppMode)
	r := gin.Default()

	/*
		前端展示页面接口
	*/
	router := r.Group("api/v1")
	{
		// 用户信息模块
		router.POST("user/add", v1.AddUser)
		router.GET("user/:id", v1.GetUserInfo)
		router.GET("users", v1.GetUsers)
	}

	err := r.Run(setting.HttpPort)
	if err != nil {
		return
	}

}
