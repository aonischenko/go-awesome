package controller

import (
	"github.com/gin-gonic/gin"
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
	div(ctx, ReadUrlParams)
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
	div(ctx, ReadBody)
}

func div(ctx *gin.Context, opReader ModelReader) {
	op := &model.BinaryOp{Operation: model.Operation{Name: "division"}}
	log := RequestLogger(ctx)

	if err := opReader(ctx, op); err != nil {
		apiError := model.NewApiError(http.StatusUnprocessableEntity, "can't read input entity", err.Error())
		log.Debugf("API Error: %s. Details: %s", apiError.Message, apiError.Details)
		ctx.JSON(apiError.Status, apiError)
		return
	}

	res, err := ops.DivWithRemainder(op.Left, op.Right)
	if err != nil {
		apiError := model.NewApiError(http.StatusBadRequest, "operation error", err.Error())
		log.Debugf("API Error: %s. Details: %s", apiError.Message, apiError.Details)
		ctx.JSON(apiError.Status, apiError)
		return
	}

	ctx.JSON(http.StatusOK, model.OpResult{
		Operation: op,
		Success:   true,
		Result:    res.AsPlain(),
	})
}