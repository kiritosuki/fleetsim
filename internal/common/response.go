package common

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Result struct {
	Code    int         `json:"code,omitempty"` // 1表示成功 2表示失败
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

func Success(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, Result{
		Code:    1,
		Message: "success",
		Data:    data,
	})
}

func Error(c *gin.Context, msg string) {
	c.JSON(http.StatusOK, Result{
		Code:    2,
		Message: msg,
	})
}
