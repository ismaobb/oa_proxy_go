package models

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Result        int         `json:"result"`
	Content       interface{} `json:"content"`
	ResultComment string      `json:"result_comment"`
}

const (
	ERROR   = -1
	SUCCESS = 0
)

func Ok(content interface{}, ctx *gin.Context) {
	ctx.JSON(http.StatusOK, Response{
		Result: SUCCESS, Content: content,
	})
}

func Fail(resultComment string, ctx *gin.Context) {
	ctx.AbortWithStatusJSON(http.StatusBadRequest, Response{
		Result: ERROR, ResultComment: resultComment,
	})
}
