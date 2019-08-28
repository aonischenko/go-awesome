package api

import (
	"context"
	"github.com/julienschmidt/httprouter"
	uuid "github.com/satori/go.uuid"
	log "github.com/sirupsen/logrus"
	"github.com/swaggo/http-swagger"
	"goawesome/config"
	_ "goawesome/docs" //required
	"net/http"
	"time"
)

func AppHandler(cfg config.Config) http.Handler {
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
	router.GET("/swagger/*path", swaggerHandle)

	return router
}

func swaggerHandle(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	swaggerHandler := httpSwagger.Handler(
		//URL pointing to API definition
		httpSwagger.URL("/swagger/doc.json"),
	)
	swaggerHandler.ServeHTTP(w, r)
}

// Request diagnostic middleware
func diagMiddleware(next httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		contextKey, headerName := "requestId", "X-Request-Id"
		requestId := r.Header.Get(headerName)
		if requestId == "" {
			requestId = uuid.NewV4().String()
		}
		w.Header().Add(headerName, requestId)
		ctx := r.WithContext(context.WithValue(r.Context(), contextKey, requestId))
		next(w, ctx, p)
	}
}

// Request logging middleware
// Simply wraps the handler function with some log messages
func logMiddleware(next httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		logger := log.WithField("requestId", r.Context().Value("requestId"))
		//contextKey := "logger"
		//r.WithContext(context.WithValue(r.Context(), contextKey, logger))
		start := time.Now()
		//todo check if we have to compare to current log lvl first
		logger.Tracef("%s %s : Request started", r.Method, r.URL.Path)
		next(w, r, p)
		logger.Tracef("%s %s : Request finished in %v", r.Method, r.URL.Path, time.Since(start))
	}
}
