package api

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"goawesome/handler"
	"goawesome/model"
	"goawesome/ops"
	"net/http"
)

/*
API V2 routes
*/
type V2 struct {
	Version string
}

func NewV2() *V2 {
	return &V2{Version: Version2}
}

func (v *V2) ListRoutes() Routes {
	return Routes{
		{Method: "GET", Path: fmt.Sprint("/", v.Version, "/div"), Handler: v.divByGet},
		{Method: "PUT", Path: fmt.Sprint("/", v.Version, "/div"), Handler: v.divByPut},
	}
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
func (v *V2) divByGet(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	div2(w, r, handler.ReadUrlParams)
}

// @Summary Division using request body
// @Description Should return status 200 with an division operation result
// @ID v2PutDiv
// @Accept json
// @Produce json
// @Success 200 {object} model.OpResult
// @Failure 400 {object} model.ApiError
// @Router /v2/div [put]
func (v *V2) divByPut(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	div2(w, r, handler.ReadBody)
}

func div2(w http.ResponseWriter, r *http.Request, f handler.ModelReader) {
	op := &model.BinaryOp{Operation: model.Operation{Name: "division"}}
	log := RequestLogger(r.Context())
	// test request logger
	log.Trace("v2 div func called")

	if err := f(r, op); err != nil {
		apiError := model.NewApiError(http.StatusBadRequest, "can't read input entity", err.Error())
		log.Debugf("Api Error: %s. Details: %s", apiError.Message, apiError.Details)
		handler.WriteError(w, apiError)
		return
	}

	res, err := ops.DivWithRemainder(op.Left, op.Right)
	if err != nil {
		log.Debugf("Api Error: %s. Details: %s", "operation error", err.Error())
		handler.WriteOk(w, model.OpResult{
			Operation: op,
			Success:   false,
			Result:    err.Error(),
		})
		return
	}

	handler.WriteOk(w, model.OpResult{
		Operation: op,
		Success:   true,
		Result:    res.AsPeriodic(),
	})
}
