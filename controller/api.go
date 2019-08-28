package controller

import "github.com/gin-gonic/gin"

type API interface {
	RegisterHandlers(r *gin.RouterGroup)
}
