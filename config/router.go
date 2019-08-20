package config

import (
	"github.com/julienschmidt/httprouter"
	"github.com/sirupsen/logrus"
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

	return router
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
