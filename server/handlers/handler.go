package handlers

import (
	"database/sql"
	"net/http"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/yafiesetyo/merempah-api-clone/helper"
	"github.com/yafiesetyo/merempah-api-clone/usecase"
)

type Handler struct {
	FiberApp   *fiber.App
	ContractUC *usecase.ContractUC
	Db         *sql.DB
}

func (h Handler) SendResponse(ctx *fiber.Ctx, data interface{}, meta interface{}, err interface{}, code int) error {
	if code == 0 && err != nil {
		code = http.StatusUnprocessableEntity
		err = err.(error).Error()
	}

	if code != http.StatusOK && err != nil {
		return h.SendErrorResponse(ctx, err, code)
	}

	if code == http.StatusAccepted && code != http.StatusOK && err != nil {
		return h.SendAcceptedResponse(ctx, data, code)
	}

	return h.SendSuccessResponse(ctx, data, meta)
}

//send response if status code 201
func (h Handler) SendAcceptedResponse(ctx *fiber.Ctx, data interface{}, meta interface{}) error {
	response := helper.SuccessResponse(data, meta)

	return ctx.Status(http.StatusAccepted).JSON(response)
}

//send response if status code 200
func (h Handler) SendSuccessResponse(ctx *fiber.Ctx, data interface{}, meta interface{}) error {
	response := helper.SuccessResponse(data, meta)

	return ctx.Status(http.StatusOK).JSON(response)
}

//send response if status code != 200
func (h Handler) SendErrorResponse(ctx *fiber.Ctx, err interface{}, code int) error {
	response := helper.ErrorResponse(err)

	return ctx.Status(code).JSON(response)
}

//send response if status code 200
func (h Handler) SendFileResponse(ctx *fiber.Ctx, data, contentType string) error {
	fileRes, err := os.OpenFile(data, os.O_RDWR, 0644)
	if err != nil {
		return h.SendErrorResponse(ctx, "a"+err.Error(), http.StatusBadRequest)
	}

	fi, err := fileRes.Stat()
	if err != nil {
		return h.SendErrorResponse(ctx, "b"+err.Error(), http.StatusBadRequest)
	}

	ctx.Set("Content-Type", contentType)
	return ctx.Status(http.StatusOK).SendStream(fileRes, int(fi.Size()))
}
