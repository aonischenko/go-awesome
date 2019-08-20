package handler

import (
	"github.com/julienschmidt/httprouter"
	"github.com/sirupsen/logrus"
	"goawesome/model"
	"goawesome/ops"
	"net/http"
)

func DivGet(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	div(w, r, ReadUrlParams)
}

func DivPut(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	div(w, r, ReadBody)
}

func div(w http.ResponseWriter, r *http.Request, f requestReader) {
	op := &model.BinaryOp{Operation: model.Operation{Name: "division"}}

	if err := f(r, op); err != nil {
		apiError := model.NewApiError(http.StatusUnprocessableEntity, "Can't read input entity", err.Error())
		WriteError(w, apiError)
		logrus.Tracef("Can't read input entity: %s", err.Error())
		return
	}

	res, err := ops.DivWithRemainder(op.Left, op.Right)
	if err != nil {
		apiError := model.NewApiError(http.StatusBadRequest, "Operation error", err.Error())
		WriteError(w, apiError)
		logrus.Tracef("Operation error: %s", err.Error())
		return
	}

	WriteOk(w, model.OpResult{
		Operation: op,
		Success:   true,
		Result:    res.AsPlain(),
	})
}
