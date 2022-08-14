package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// 聊天页面
func Index(c *gin.Context) {

	data := gin.H{
		"title":        "聊天首页",
		"helloMessage": "hello osman",
	}
	c.HTML(http.StatusOK, "index.tpl", data)
}
