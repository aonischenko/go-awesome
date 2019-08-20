package handler

import (
	"github.com/julienschmidt/httprouter"
	"github.com/sirupsen/logrus"
	"goawesome/model"
	"goawesome/ops"
	"net/http"
)

func DivGetV2(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	div2(w, r, ReadUrlParams)
}

func DivPutV2(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	div2(w, r, ReadBody)
}

func div2(w http.ResponseWriter, r *http.Request, f requestReader) {
	op := &model.BinaryOp{Operation: model.Operation{Name: "division"}}

	if err := f(r, op); err != nil {
		apiError := model.NewApiError(http.StatusBadRequest, "Can't read input entity", err.Error())
		WriteError(w, apiError)
		logrus.Tracef("Can't read input entity: %s", err.Error())
		return
	}

	res, err := ops.DivWithRemainder(op.Left, op.Right)
	if err != nil {
		WriteOk(w, model.OpResult{
			Operation: op,
			Success:   false,
			Result:    err.Error(),
		})
		logrus.Tracef("Operation error: %s", err.Error())
		return
	}

	WriteOk(w, model.OpResult{
		Operation: op,
		Success:   true,
		Result:    res.AsPeriodic(),
	})
}
