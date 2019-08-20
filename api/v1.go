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
		v1GetDiv(),
		v1PutDiv(),
	}
	return routes
}

// @Summary Division using request url params
// @Description Should return status 200 with an division operation result
// @ID v1GetDiv
// @Produce json
// @Success 200 {object} model.OpResult
// @Failure 400 {object} model.ApiError
// @Failure 422 {object} model.ApiError
// @Param x query int true "division operation numerator"
// @Param y query int true "division operation denominator"
// @Router /v1/div?x={x}&y={y} [get]
func v1GetDiv() model.Route {
	return model.Route{
		Name:        "Division using request url params",
		Method:      "GET",
		Path:        "/div",
		HandlerFunc: handler.DivGet,
	}
}

// @Summary Division using request body
// @Description Should return status 200 with an division operation result
// @ID v1PutDiv
// @Accept json
// @Produce json
// @Success 200 {object} model.OpResult
// @Failure 400 {object} model.ApiError
// @Failure 422 {object} model.ApiError
// @Router /v1/div [put]
func v1PutDiv() model.Route {
	return model.Route{
		Name:        "Division using request body",
		Method:      "PUT",
		Path:        "/div",
		HandlerFunc: handler.DivPut,
	}
}
