package api

import (
	"goawesome/handler"
	"goawesome/model"
)

/*
API V2 routes
*/
func v2() model.Routes {
	routes := model.Routes{
		{
			Name:        "Division using request url params",
			Method:      "GET",
			Path:        "/div",
			HandlerFunc: handler.DivGetV2,
		},
		{
			Name:        "Division using request body",
			Method:      "PUT",
			Path:        "/div",
			HandlerFunc: handler.DivPutV2,
		},
	}
	return routes
}
