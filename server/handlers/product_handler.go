package handlers

import (
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
	request "github.com/yafiesetyo/merempah-api-clone/server/requests"
	"github.com/yafiesetyo/merempah-api-clone/usecase"
)

// ProductHandler...
type ProductHandler struct {
	Handler
}

// SelectAll...
func (h *ProductHandler) SelectAll(ctx *fiber.Ctx) error {
	keyword := ctx.Query("keyword")

	productUc := usecase.ProductUC{ContractUC: h.ContractUC}
	res, err := productUc.SelectAll(keyword)
	return h.SendResponse(ctx, res, nil, err, 0)
}

// FindById
func (h *ProductHandler) FindById(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	convert, err := strconv.Atoi(id)
	if err != nil {
		return h.SendResponse(ctx, nil, nil, err, http.StatusBadRequest)
	}
	productUc := usecase.ProductUC{ContractUC: h.ContractUC}
	res, err := productUc.FindById(int64(convert))
	return h.SendResponse(ctx, res, nil, err, 0)
}

// Add
func (h *ProductHandler) Add(ctx *fiber.Ctx) error {
	input := new(request.ProductRequest)
	if err := ctx.BodyParser(input); err != nil {
		return h.SendResponse(ctx, nil, nil, err, http.StatusBadRequest)
	}
	productUc := usecase.ProductUC{ContractUC: h.ContractUC}
	res, err := productUc.Add(input)
	return h.SendResponse(ctx, res, nil, err, 0)
}

// Edit...
func (h *ProductHandler) Edit(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	if id == "" {
		return h.SendResponse(ctx, nil, nil, nil, http.StatusBadRequest)
	}

	convert, err := strconv.Atoi(id)
	if err != nil {
		return h.SendResponse(ctx, nil, nil, err, http.StatusBadRequest)
	}

	input := new(request.ProductRequest)
	if err := ctx.BodyParser(input); err != nil {
		return h.SendResponse(ctx, nil, nil, err, http.StatusBadRequest)
	}

	productUc := usecase.ProductUC{ContractUC: h.ContractUC}
	res, err := productUc.Edit(int64(convert), input)
	return h.SendResponse(ctx, res, nil, err, 0)
}

// Delete
func (h *ProductHandler) Delete(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	if id == "" {
		return h.SendResponse(ctx, nil, nil, nil, http.StatusBadRequest)
	}

	convert, err := strconv.Atoi(id)
	if err != nil {
		return h.SendResponse(ctx, nil, nil, err, http.StatusBadRequest)
	}

	productUc := usecase.ProductUC{ContractUC: h.ContractUC}
	res, err := productUc.Delete(int64(convert))
	return h.SendResponse(ctx, res, nil, err, 0)
}
