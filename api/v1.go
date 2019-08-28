package api

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	log "github.com/sirupsen/logrus"
	"goawesome/handler"
	"goawesome/model"
	"goawesome/ops"
	"net/http"
)

type V1 struct {
	Version string
}

func NewV1() *V1 {
	return &V1{Version: Version1}
}

/*
API V1 routes
*/
func (v *V1) ListRoutes() Routes {
	return Routes{
		{Method: "GET", Path: fmt.Sprintf("/%s/div", v.Version), Handle: v.divByGet},
		{Method: "PUT", Path: fmt.Sprintf("/%s/div", v.Version), Handle: v.divByPut},
	}
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
func (v *V1) divByGet(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	div(w, r, handler.ReadUrlParams)
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
func (v *V1) divByPut(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	div(w, r, handler.ReadBody)
}

func div(w http.ResponseWriter, r *http.Request, f handler.RequestReader) {
	op := &model.BinaryOp{Operation: model.Operation{Name: "division"}}

	if err := f(r, op); err != nil {
		apiError := model.NewApiError(http.StatusUnprocessableEntity, "can't read input entity", err.Error())
		log.Debugf("Api Error: %s. Details: %s", apiError.Message, apiError.Details)
		handler.WriteError(w, apiError)
		return
	}

	res, err := ops.DivWithRemainder(op.Left, op.Right)
	if err != nil {
		apiError := model.NewApiError(http.StatusBadRequest, "operation error", err.Error())
		log.Debugf("Api Error: %s. Details: %s", apiError.Message, apiError.Details)
		handler.WriteError(w, apiError)
		return
	}

	handler.WriteOk(w, model.OpResult{
		Operation: op,
		Success:   true,
		Result:    res.AsPlain(),
	})
}
