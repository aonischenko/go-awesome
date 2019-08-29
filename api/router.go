package api

import (
	"context"
	"github.com/julienschmidt/httprouter"
	uuid "github.com/satori/go.uuid"
	"github.com/sirupsen/logrus"
	"github.com/swaggo/http-swagger"
	. "goawesome/config"
	_ "goawesome/docs" //required
	"net/http"
	"time"
)

const (
	ContextLoggerKey    = "logger"
	ContextRequestIdKey = "requestId"
	HeaderRequestIdKey  = "X-Request-Id"
)

func AppHandler(cfg Config) http.Handler {
	router := httprouter.New()

	for _, v := range Versions() {
		for _, h := range v.ListRoutes() {
			var handler httprouter.Handle
			handler = h.Handler
			handler = logMiddleware(handler)
			handler = diagMiddleware(handler)

			router.Handle(h.Method, h.Path, handler)
		}
	}

	//adding swagger handlers
	if cfg.SwagEnable {
		router.GET("/swagger/*path", swaggerHandler)
	}

	return router
}

func swaggerHandler(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	swaggerHandler := httpSwagger.Handler(
		//URL pointing to API definition
		httpSwagger.URL("/swagger/doc.json"),
	)
	swaggerHandler.ServeHTTP(w, r)
}

// Request diagnostic middleware
func diagMiddleware(next httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		requestId := r.Header.Get(HeaderRequestIdKey)
		if requestId == "" {
			requestId = uuid.NewV4().String()
		}
		w.Header().Add(HeaderRequestIdKey, requestId)
		request := r.WithContext(context.WithValue(r.Context(), ContextRequestIdKey, requestId))
		next(w, request, p)
	}
}

// Request logging middleware
// Simply wraps the handler function with some log messages
func logMiddleware(next httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		logger := Log.WithField("requestId", r.Context().Value(ContextRequestIdKey))
		request := r.WithContext(context.WithValue(r.Context(), ContextLoggerKey, logger))
		start := time.Now()
		//todo check if we have to compare to current log lvl first
		logger.Tracef("%s %s : Request started", r.Method, r.URL.Path)
		next(w, request, p)
		logger.Tracef("%s %s : Request finished in %v", r.Method, r.URL.Path, time.Since(start))
	}
}

func RequestLogger(c context.Context) logrus.Entry {
	if contextLogger := c.Value(ContextLoggerKey); contextLogger != nil {
		if logger, ok := contextLogger.(*logrus.Entry); ok {
			return *logger
		}
	}
	return Log
}
