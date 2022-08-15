package webhandler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// 聊天页面
func IndexPage(c *gin.Context) {

	data := gin.H{
		"title":        "Go Live Chat",
		"helloMessage": "hello osman",
	}
	c.HTML(http.StatusOK, "index.tpl", data)
}
