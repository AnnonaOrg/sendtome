package api_handler

import (
	"github.com/AnnonaOrg/pkg/errno"
	"github.com/umfaka/sendtome/handler"
	"github.com/umfaka/sendtome/internal/constvar"
	"github.com/gin-gonic/gin"
)

// 404 Not found
func ApiNotFound(c *gin.Context) {
	handler.SendResultResponse(c, errno.Err404, constvar.APPDesc404())
}

// API Hello
func ApiHello(c *gin.Context) {
	handler.SendResultResponse(c, errno.SayHello, constvar.APPDesc())
}

// ping
func ApiPing(c *gin.Context) {
	handler.SendResultResponse(c, errno.PONG, constvar.APPVersion())
}
