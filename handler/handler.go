package handler

import (
	"errors"
	"projeto-teste/fail"
	"projeto-teste/model"
	"projeto-teste/request"
	"projeto-teste/response"
	"projeto-teste/service"

	"github.com/gin-gonic/gin"
)

type Handler[RS response.IResponse, M model.IModel[RS], CRQ request.IRequest[M]] struct {
	Service *service.Service[RS, M, CRQ]
}

func InitHandler[RS response.IResponse, M model.IModel[RS], CRQ request.IRequest[M]](service *service.Service[RS, M, CRQ]) Handler[RS, M, CRQ] {
	return Handler[RS, M, CRQ]{Service: service}
}

func (h *Handler[RS, M, CRQ]) Create(ctx *gin.Context) {
	var request CRQ

	if err := ctx.ShouldBindJSON(&request); err != nil {
		handleError(ctx, fail.ErrorToBadRequest(err))
		return
	}

	response, err := h.Service.Create(request)

	if err != nil {
		handleError(ctx, err)
		return
	}

	sendObject(ctx, 201, response)
}

var (
	badRequest          fail.BadRequest
	internalServerError fail.InternalServerError
)

func sendObject(context *gin.Context, statusCode int, object any) {
	context.JSON(statusCode, object)
}

func handleError(context *gin.Context, err error) bool {

	if err == nil {
		return false
	}

	message := func(status int, message string) {
		context.JSON(status, gin.H{
			"http_status":   status,
			"error_message": message,
		})
	}

	if errors.As(err, &badRequest) {
		message(400, err.Error())
		return true
	}

	if errors.As(err, &internalServerError) {
		message(500, err.Error())
		return true
	}

	return true
}

func sendSuccess(context *gin.Context, statusCode int) {
	context.JSON(statusCode, gin.H{
		"message": "operation successfully completed!",
	})
}
