package controller

import "github.com/gin-gonic/gin"

type ModelReader func(ctx *gin.Context, model interface{}) error

func ReadBody(ctx *gin.Context, model interface{}) error {
	return ctx.ShouldBindJSON(model)
}

func ReadUrlParams(ctx *gin.Context, model interface{}) error {
	return ctx.ShouldBindQuery(model)
}
