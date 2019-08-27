package config

import (
	"context"
	"github.com/julienschmidt/httprouter"
	uuid "github.com/satori/go.uuid"
	log "github.com/sirupsen/logrus"
	"github.com/swaggo/http-swagger"
	"github.com/urfave/negroni"
	"goawesome/api"
	_ "goawesome/docs" //required
	"net/http"
	"time"
)

func AppHandler() http.Handler {
	n := negroni.New()
	n.Use(negroni.HandlerFunc(diagMiddleware))
	n.Use(negroni.HandlerFunc(logMiddleware))
	n.Use(negroni.NewRecovery())

	router := httprouter.New()
	n.UseHandler(router)
	for _, v := range api.Versions() {
		v.RegisterHandlers(router)
	}

	//adding swagger handlers
	//todo move swagger handlers into new group w|o additional middleware
	router.GET("/swagger/*path", swaggerHandle)

	return n
}

func swaggerHandle(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	swaggerHandler := httpSwagger.Handler(
		//URL pointing to API definition"
		httpSwagger.URL("/swagger/doc.json"),
	)
	swaggerHandler.ServeHTTP(w, r)
}

// Request diagnostic middleware
func diagMiddleware(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	contextKey, headerName := "requestId", "X-Request-Id"
	requestId := r.Header.Get(headerName)
	if requestId == "" {
		requestId = uuid.NewV4().String()
	}
	w.Header().Add(headerName, requestId)
	ctx := r.WithContext(context.WithValue(r.Context(), contextKey, requestId))
	next(w, ctx)
}

// Request logging middleware
// Simply wraps the handler function with some log messages
func logMiddleware(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	logger := log.WithField("requestId", r.Context().Value("requestId"))
	//contextKey := "logger"
	//r.WithContext(context.WithValue(r.Context(), contextKey, logger))
	start := time.Now()
	//todo check if we have to compare to current log lvl first
	logger.Tracef("%s %s : Request started", r.Method, r.URL.Path)
	next(w, r)
	logger.Tracef("%s %s : Request finished in %v", r.Method, r.URL.Path, time.Since(start))
}
