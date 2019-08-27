package handler

import (
	"context"
	"fmt"
	uuid "github.com/satori/go.uuid"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/negroni"
	"net/http"
	"time"
)

// Request diagnostic middleware
func DiagMiddleware(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
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
func LogMiddleware(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	logger := log.WithFields(
		log.Fields{
			"request":   fmt.Sprintf("%s %s", r.Method, r.URL.Path),
			"requestId": r.Context().Value("requestId"),
		})

	//contextKey := "logger"
	//r.WithContext(context.WithValue(r.Context(), contextKey, logger))

	start := time.Now()
	logger.Trace("Request started")
	rw := negroni.NewResponseWriter(w)

	next(rw, r)

	logger = logger.WithField("responseStatus", rw.Status())
	logger.Tracef("Request finished in %v", time.Since(start))
}
