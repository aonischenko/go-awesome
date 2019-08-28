package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"goawesome/handler"
	"goawesome/model"
	"goawesome/ops"
	"net/http"
)

/*
API V1 routes
*/
type V1 struct {
}

func (v V1) RegisterHandlers(r *gin.RouterGroup) {
	r.GET("/div", v.divByGet)
	r.PUT("/div", v.divByPut)
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
func (v *V1) divByGet(ctx *gin.Context) {
	div(ctx, handler.ReadUrlParams)
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
func (v *V1) divByPut(ctx *gin.Context) {
	div(ctx, handler.ReadBody)
}

func div(ctx *gin.Context, f handler.RequestReader) {
	op := &model.BinaryOp{Operation: model.Operation{Name: "division"}}

	if err := f(ctx.Request, op); err != nil {
		apiError := model.NewApiError(http.StatusUnprocessableEntity, "can't read input entity", err.Error())
		logrus.Debugf("Api Error: %s. Details: %s", apiError.Message, apiError.Details)
		handler.WriteError(ctx.Writer, apiError)
		return
	}

	res, err := ops.DivWithRemainder(op.Left, op.Right)
	if err != nil {
		apiError := model.NewApiError(http.StatusBadRequest, "operation error", err.Error())
		logrus.Debugf("Api Error: %s. Details: %s", apiError.Message, apiError.Details)
		handler.WriteError(ctx.Writer, apiError)
		return
	}

	handler.WriteOk(ctx.Writer, model.OpResult{
		Operation: op,
		Success:   true,
		Result:    res.AsPlain(),
	})
}