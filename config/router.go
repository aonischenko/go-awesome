package config

import (
	"github.com/julienschmidt/httprouter"
	"github.com/sirupsen/logrus"
	"github.com/swaggo/http-swagger"
	_ "goawesome/docs" //required
	"goawesome/model"
	"net/http"
	"time"
)

func AppRouter(versions model.Versions) *httprouter.Router {
	router := httprouter.New()
	for _, version := range versions {
		prefix := "/" + version.Version
		for _, route := range version.Routes {
			var handle httprouter.Handle

			handle = route.HandlerFunc
			handle = logRequest(handle)

			router.Handle(route.Method, prefix+route.Path, handle)
		}
	}

	//adding swagger handlers
	router.GET("/swagger/*path", swaggerHandler())

	return router
}

func swaggerHandler() httprouter.Handle {
	handler := httpSwagger.Handler(
		httpSwagger.URL("/swagger/doc.json"), //The url pointing to API definition"
	)
	return func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
		handler.ServeHTTP(w, r)
	}
}

// A Logger function which simply wraps the handler function around some log messages
func logRequest(fn httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, param httprouter.Params) {
		start := time.Now()
		//todo check if we have to compare to current log lvl first
		logrus.Debugf("%s %s : Request started", r.Method, r.URL.Path)
		fn(w, r, param)
		logrus.Debugf("%s %s : Request finished in %v", r.Method, r.URL.Path, time.Since(start))
	}
}
