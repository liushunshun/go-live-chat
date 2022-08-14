package config

import (
	"go-live-chat/server/handlers"

	"github.com/gin-gonic/gin"
)

func InitWebRouters(router *gin.Engine) {
	router.LoadHTMLGlob("views/*")

	// index
	homeRouter := router.Group("")
	{
		homeRouter.GET("/", handlers.Index)
	}

	// 用户组
	// userRouter := router.Group("/user")
	// {
	// 	userRouter.GET("/list", user.List)
	// 	userRouter.GET("/online", user.Online)
	// 	userRouter.POST("/sendMessage", user.SendMessage)
	// 	userRouter.POST("/sendMessageAll", user.SendMessageAll)
	// }

	// router.POST("/user/online", user.Online)
}
