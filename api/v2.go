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
		v2GetDiv(),
		v2PutDiv(),
	}
	return routes
}

// @Summary Division using request url params
// @Description Should return status 200 with an division operation result
// @ID v2GetDiv
// @Produce json
// @Success 200 {object} model.OpResult
// @Failure 400 {object} model.ApiError
// @Param x query int true "division operation numerator"
// @Param y query int true "division operation denominator"
// @Router /v2/div?x={x}&y={y} [get]
func v2GetDiv() model.Route {
	return model.Route{
		Name:        "Division using request url params",
		Method:      "GET",
		Path:        "/div",
		HandlerFunc: handler.DivGet,
	}
}

// @Summary Division using request body
// @Description Should return status 200 with an division operation result
// @ID v2PutDiv
// @Accept json
// @Produce json
// @Success 200 {object} model.OpResult
// @Failure 400 {object} model.ApiError
// @Router /v2/div [put]
func v2PutDiv() model.Route {
	return model.Route{
		Name:        "Division using request body",
		Method:      "PUT",
		Path:        "/div",
		HandlerFunc: handler.DivPut,
	}
}
