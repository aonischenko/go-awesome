package controller

import (
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	log "github.com/sirupsen/logrus"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
	"goawesome/config"
	_ "goawesome/docs" //required
	"net/http"
	"time"
)

type API interface {
	RegisterHandlers(r *gin.RouterGroup)
}

func AppHandler(cfg config.Config) http.Handler {
	r := gin.New()

	g1 := r.Group("/v1")
	g1.Use(diagMiddleware, logMiddleware)
	v1 := V1{}
	v1.RegisterHandlers(g1)

	g2 := r.Group("/v2")
	g2.Use(gin.Recovery(), diagMiddleware, logMiddleware)
	v2 := V2{}
	v2.RegisterHandlers(g2)

	if cfg.SwagEnable {
		r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}

	return r
}

// Request diagnostic middleware
func diagMiddleware(ctx *gin.Context) {
	contextKey, headerName := "requestId", "X-Request-Id"
	requestId := ctx.Request.Header.Get(headerName)
	if requestId == "" {
		requestId = uuid.NewV4().String()
	}
	ctx.Writer.Header().Add(headerName, requestId)
	ctx.Set(contextKey, requestId)
}

// Request logging middleware
// Simply wraps the handler function with some log messages
func logMiddleware(ctx *gin.Context) {
	logger := log.WithField("requestId", ctx.Value("requestId"))
	//contextKey := "logger"
	//r.WithContext(context.WithValue(r.Context(), contextKey, logger))
	start := time.Now()
	//todo check if we have to compare to current log lvl first
	logger.Tracef("%s %s : Request started", ctx.Request.Method, ctx.Request.URL.Path)
	ctx.Next()
	logger.Tracef("%s %s : Request finished in %v", ctx.Request.Method, ctx.Request.URL.Path, time.Since(start))
}
