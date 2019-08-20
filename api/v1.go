package api

import (
	"goawesome/handler"
	"goawesome/model"
)

/*
API V1 routes
*/
func v1() model.Routes {
	routes := model.Routes{
		{
			Name:        "Division using request url params",
			Method:      "GET",
			Path:        "/div",
			HandlerFunc: handler.DivGet,
		},
		{
			Name:        "Division using request body",
			Method:      "PUT",
			Path:        "/div",
			HandlerFunc: handler.DivPut,
		},
	}
	return routes
}
