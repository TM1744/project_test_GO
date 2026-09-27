package handler

import (
	"errors"
	"projeto-teste/converter"
	"projeto-teste/fail"
	"projeto-teste/model"
	"projeto-teste/request"
	"projeto-teste/response"
	"projeto-teste/service"

	"github.com/gin-gonic/gin"
)

type Handler[Model model.IModel, CRequest request.IRequest[Model], Response response.IResponse[Model]] struct {
	Service     *service.Service[Model, CRequest, Response]
	newCRequest func() CRequest
}

func InitHandler[Model model.IModel, CRequest request.IRequest[Model], Response response.IResponse[Model]](
		service *service.Service[Model, CRequest, Response],
		newCRequest func () CRequest,
	) *Handler[Model, CRequest, Response] {
	return &Handler[Model, CRequest, Response]{
		Service: service,
		newCRequest: newCRequest,
	}
}

func (h *Handler[Model, CRequest, Response]) Create(ctx *gin.Context) {
	request := h.newCRequest()

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

func (h *Handler[Model, CRequest, Response]) Read(ctx *gin.Context) {
	var err error
	
	id, err := catchID(ctx)

	if err != nil {
		handleError(ctx, err)
		return
	}

	response, err := h.Service.Read(id)

	if err != nil {
		handleError(ctx, err)
		return
	}

	sendObject(ctx, 200, response)
}

func (h *Handler[Model, CRequest, Response]) Delete(ctx *gin.Context) {
	var err error

	id, err := catchID(ctx)

	if err != nil {
		handleError(ctx, err)
		return
	}

	err = h.Service.Delete(id)

	if err != nil {
		handleError(ctx, err)
		return
	}

	sendSuccess(ctx, 200)
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

func catchID(ctx *gin.Context) (*int64, error) {
	id := ctx.Param("id")
	number, err := converter.StringToInt64(id)
	return &number, err
}
