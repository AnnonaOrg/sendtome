package handler

import (
	"net/http"

	"github.com/AnnonaOrg/pkg/errno"
	"github.com/gin-gonic/gin"
	"github.com/umfaka/sendtome/core/response"
)

func SendResponse(c *gin.Context, err error, data interface{}) {
	code, message := errno.DecodeErr(err)

	c.JSON(http.StatusOK, response.Response{
		Code:    code,
		Message: message,
		Data:    data,
	})

}

func SendResultResponse(c *gin.Context, err error, data interface{}) {
	code, message := errno.DecodeErr(err)
	c.JSON(http.StatusOK, response.Response{
		Code:    code,
		Message: message,
		Data: response.ResultResponse{
			Result: data,
		},
	})
}

func SendRedirect(c *gin.Context, data string) {
	c.Redirect(http.StatusMovedPermanently, data)
}

func SendRedirect302(c *gin.Context, data string) {
	c.Redirect(http.StatusFound, data)
}
