package controller

import (
	"backend/consts"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 全局统一返回结构体
type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func Success(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, Response{
		Code: consts.CodeSuccess,
		Msg:  consts.MsgSuccess,
		Data: data,
	})
}

func Failed(c *gin.Context, code int, msg string) {
	c.JSON(http.StatusOK, Response{
		Code: code,
		Msg:  msg,
		Data: struct{}{},
	})
}
