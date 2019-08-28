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
API V2 routes
*/
type V2 struct {
}

func (v V2) RegisterHandlers(r *gin.RouterGroup) {
	r.GET("/div", v.divByGet)
	r.PUT("/div", v.divByPut)
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
func (v *V2) divByGet(ctx *gin.Context) {
	div2(ctx, handler.ReadUrlParams)
}

// @Summary Division using request body
// @Description Should return status 200 with an division operation result
// @ID v2PutDiv
// @Accept json
// @Produce json
// @Success 200 {object} model.OpResult
// @Failure 400 {object} model.ApiError
// @Router /v2/div [put]
func (v *V2) divByPut(ctx *gin.Context) {
	div2(ctx, handler.ReadBody)
}

func div2(ctx *gin.Context, f handler.RequestReader) {
	op := &model.BinaryOp{Operation: model.Operation{Name: "division"}}

	if err := f(ctx.Request, op); err != nil {
		apiError := model.NewApiError(http.StatusBadRequest, "can't read input entity", err.Error())
		logrus.Debugf("Api Error: %s. Details: %s", apiError.Message, apiError.Details)
		handler.WriteError(ctx.Writer, apiError)
		return
	}

	res, err := ops.DivWithRemainder(op.Left, op.Right)
	if err != nil {
		logrus.Debugf("Api Error: %s. Details: %s", "operation error", err.Error())
		handler.WriteOk(ctx.Writer, model.OpResult{
			Operation: op,
			Success:   false,
			Result:    err.Error(),
		})
		return
	}

	handler.WriteOk(ctx.Writer, model.OpResult{
		Operation: op,
		Success:   true,
		Result:    res.AsPeriodic(),
	})
}
