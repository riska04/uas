package controller

import (
	"materi/helper"
	"materi/model/request"
	"materi/service"
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type NotulenController interface {
	All(context *fiber.Ctx) error
	FindByID(context *fiber.Ctx) error
	Create(context *fiber.Ctx) error
	Update(context *fiber.Ctx) error
	Delete(context *fiber.Ctx) error
	Upload(context *fiber.Ctx) error
}

type notulenController struct {
	notulenService service.NotulenService
}

// All implements NotulenController.
func (c *notulenController) All(context *fiber.Ctx) error {
	notulen, err := c.notulenService.All(context.Context())

	if err != nil {
		response := helper.APIResponse("Failed to get data", http.StatusBadRequest, "error", err.Error())
		return context.Status(http.StatusBadRequest).JSON(response)
	}
	response := helper.APIResponse("Success to get data", http.StatusOK, "success", notulen)
	return context.Status(http.StatusOK).JSON(response)
}

// Create implements NotulenController.
func (c *notulenController) Create(context *fiber.Ctx) error {
	var inputData request.NotulenCreate
	if err := context.BodyParser(&inputData); err != nil {
		response := helper.APIResponse("Failed to post data", http.StatusBadRequest, "error", err.Error())
		return context.Status(http.StatusBadRequest).JSON(response)
	}
	errors := request.ValidateStruct(&inputData)
	if errors != nil {
		response := helper.APIResponse("Failed to post data", http.StatusBadRequest, "error", errors)
		return context.Status(http.StatusBadRequest).JSON(response)
	}

	_, err := c.notulenService.Create(context.Context(), inputData)
	if err != nil {
		response := helper.APIResponse("Failed to post data", http.StatusBadRequest, "error", err.Error())
		return context.Status(http.StatusBadRequest).JSON(response)
	}
	response := helper.APIResponse("Success to post data", http.StatusOK, "success", nil)
	return context.Status(http.StatusOK).JSON(response)
}

// Delete implements NotulenController.
func (c *notulenController) Delete(context *fiber.Ctx) error {
	id, err := strconv.Atoi(context.Params("id_notulen"))
	if err != nil {
		response := helper.APIResponse("Failed to get data", http.StatusBadRequest, "error", err.Error())
		return context.Status(http.StatusBadRequest).JSON(response)

	}
	c.notulenService.Delete(context.Context(), strconv.Itoa(id))
	response := helper.APIResponse("Success to delete data", http.StatusOK, "success", nil)
	return context.Status(http.StatusOK).JSON(response)
}

// FindByID implements NotulenController.
func (c *notulenController) FindByID(context *fiber.Ctx) error {
	id, err := strconv.Atoi(context.Params("id_notulen"))
	if err != nil {
		response := helper.APIResponse("Failed to get data", http.StatusBadRequest, "error", err.Error())
		return context.Status(http.StatusBadRequest).JSON(response)

	}
	notulen, err := c.notulenService.FindByID(context.Context(), strconv.Itoa(id))

	if err != nil {
		response := helper.APIResponse("Failed to get data", http.StatusBadRequest, "error", err.Error())
		return context.Status(http.StatusBadRequest).JSON(response)
	}

	response := helper.APIResponse("Success to get data", http.StatusOK, "success", notulen)
	return context.Status(http.StatusOK).JSON(response)
}

// Update implements NotulenController.
func (c *notulenController) Update(context *fiber.Ctx) error {
	id, err := strconv.Atoi(context.Params("id_notulen"))
	if err != nil {
		response := helper.APIResponse("Failed to put data", http.StatusBadRequest, "error", err.Error())
		return context.Status(http.StatusBadRequest).JSON(response)
	}
	var inputData request.NotulenUpdate

	if err := context.BodyParser(&inputData); err != nil {
		response := helper.APIResponse("Failed to put data", http.StatusBadRequest, "error", err.Error())
		return context.Status(http.StatusBadRequest).JSON(response)
	}
	inputData.IdNotulen = strconv.Itoa(id)

	errors := request.ValidateStruct(&inputData)
	if errors != nil {
		response := helper.APIResponse("Failed to put data", http.StatusBadRequest, "error", errors)
		return context.Status(http.StatusBadRequest).JSON(response)
	}

	if err := c.notulenService.Update(context.Context(), inputData); err != nil {
		response := helper.APIResponse("Failed to put data", http.StatusBadRequest, "error", err.Error())
		return context.Status(http.StatusBadRequest).JSON(response)
	}
	response := helper.APIResponse("Success to put data", http.StatusOK, "success", nil)
	return context.Status(http.StatusOK).JSON(response)
}

// Upload implements NotulenController.
func (*notulenController) Upload(context *fiber.Ctx) error {
	panic("unimplemented")
}

func NewNotulenController(notulenService service.NotulenService) NotulenController {
	return &notulenController{
		notulenService: notulenService,
	}
}
