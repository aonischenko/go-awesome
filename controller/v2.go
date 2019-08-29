package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
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
	r.GET("/panic/:reason", v.startPanic)
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
	div2(ctx, ReadUrlParams)
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
	div2(ctx, ReadBody)
}

// @Summary Panic endpoint
// @Description Should return status 500 with an division operation result
// @ID v2StartPanic
// @Produce json
// @Success 200 {object} model.OpResult
// @Failure 500 {object}
// @Router /v2/div [put]
func (v *V2) startPanic(ctx *gin.Context) {
	reason := ops.Panic(ctx.Params.ByName("reason"))
	ctx.JSON(http.StatusOK, struct{ Reason string }{
		Reason: reason,
	})
}

func div2(ctx *gin.Context, opReader ModelReader) {
	op := &model.BinaryOp{Operation: model.Operation{Name: "division"}}
	log := RequestLogger(r.Context())
	// test request logger
	log.Trace("v2 div func called")

	if err := opReader(ctx, op); err != nil {
		apiError := model.NewApiError(http.StatusBadRequest, "can't read input entity", err.Error())
		log.Debugf("API Error: %s. Details: %s", apiError.Message, apiError.Details)
		ctx.JSON(apiError.Status, apiError)
		return
	}

	res, err := ops.DivWithRemainder(op.Left, op.Right)
	if err != nil {
		log.Debugf("API Error: %s. Details: %s", "operation error", err.Error())
		ctx.JSON(http.StatusOK, model.OpResult{
			Operation: op,
			Success:   false,
			Result:    err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, model.OpResult{
		Operation: op,
		Success:   true,
		Result:    res.AsPeriodic(),
	})
}
