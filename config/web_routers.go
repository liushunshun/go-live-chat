package config

import (
	handlers "go-live-chat/internal/webhandler"

	"github.com/gin-gonic/gin"
)

func InitWebRouters(router *gin.Engine) {
	router.LoadHTMLGlob("views/*")

	// index
	homeRouter := router.Group("")
	{
		homeRouter.GET("/", handlers.IndexPage)
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
