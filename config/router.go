package config

import (
	"github.com/julienschmidt/httprouter"
	"github.com/swaggo/http-swagger"
	"github.com/urfave/negroni"
	"goawesome/api"
	_ "goawesome/docs" //required
	"goawesome/handler"
	"net/http"
)

func AppHandler(cfg Config) http.Handler {
	router := httprouter.New()
	router.GET("/swagger/*path", swaggerHandler)

	n := negroni.New()
	n.Use(negroni.NewRecovery())
	n.UseHandler(router)

	// add middleware for a specific route and get params from route
	nApi := negroni.New(
		negroni.HandlerFunc(handler.DiagMiddleware),
		negroni.HandlerFunc(handler.LogMiddleware),
	)
	for _, v := range api.Apis() {
		for _, route := range v.Routes() {
			h := nApi.With(negroni.WrapFunc(withParams(router, route.Handle)))
			router.Handler(route.Method, route.Path, h)
		}
	}

	return n

	//n := negroni.New()
	//n.Use(negroni.HandlerFunc(handler.DiagMiddleware))
	//n.Use(negroni.HandlerFunc(handler.LogMiddleware))
	//n.Use(negroni.NewRecovery())
	//
	//router := httprouter.New()
	//router.Handler(n.With(negroni.Wrap(router)))
	//
	//for _, v := range api.Apis() {
	//	v.RegisterHandlers(router)
	//}
	//n.UseHandler(router)
	//
	//subRouter := httprouter.New()
	//subRouter.Handle("GET", "/swagger/*", swaggerHandler())
	//
	////adding swagger handlers
	////todo move swagger handlers into new group w|o additional middleware
	////router.GET("/swagger/*path", swaggerHandler())
	//
	//n.UseHandler(subRouter)
	//return n
}

func swaggerHandler(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	swaggerHandler := httpSwagger.Handler(
		httpSwagger.URL("/swagger/doc.json"), //The url pointing to API definition"
	)
	swaggerHandler.ServeHTTP(w, r)
}

// helper function to call controller from middleware having access to URL params
func withParams(router *httprouter.Router, handler httprouter.Handle) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, params, _ := router.Lookup(r.Method, r.URL.Path)
		handler(w, r, params)
	}
}
