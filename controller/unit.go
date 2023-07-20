package controller

import (
	"materi/helper"
	"materi/model/request"
	"materi/service"
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type UnitController interface {
	All(context *fiber.Ctx) error
	FindByID(context *fiber.Ctx) error
	Create(context *fiber.Ctx) error
	Update(context *fiber.Ctx) error
	Delete(context *fiber.Ctx) error
	Upload(context *fiber.Ctx) error
}

type unitController struct {
	unitService service.UnitService
}

// All implements UnitController.
func (c *unitController) All(context *fiber.Ctx) error {
	unit, err := c.unitService.All(context.Context())

	if err != nil {
		response := helper.APIResponse("Failed to get data", http.StatusBadRequest, "error", err.Error())
		return context.Status(http.StatusBadRequest).JSON(response)
	}
	response := helper.APIResponse("Success to get data", http.StatusOK, "success", unit)
	return context.Status(http.StatusOK).JSON(response)
}

// Create implements UnitController.
func (c *unitController) Create(context *fiber.Ctx) error {
	var inputData request.UnitCreate
	if err := context.BodyParser(&inputData); err != nil {
		response := helper.APIResponse("Failed to post data", http.StatusBadRequest, "error", err.Error())
		return context.Status(http.StatusBadRequest).JSON(response)
	}
	errors := request.ValidateStruct(&inputData)
	if errors != nil {
		response := helper.APIResponse("Failed to post data", http.StatusBadRequest, "error", errors)
		return context.Status(http.StatusBadRequest).JSON(response)
	}

	_, err := c.unitService.Create(context.Context(), inputData)
	if err != nil {
		response := helper.APIResponse("Failed to post data", http.StatusBadRequest, "error", err.Error())
		return context.Status(http.StatusBadRequest).JSON(response)
	}
	response := helper.APIResponse("Success to post data", http.StatusOK, "success", nil)
	return context.Status(http.StatusOK).JSON(response)
}

// Delete implements UnitController.
func (c *unitController) Delete(context *fiber.Ctx) error {
	id, err := strconv.Atoi(context.Params("id_unit"))
	if err != nil {
		response := helper.APIResponse("Failed to get data", http.StatusBadRequest, "error", err.Error())
		return context.Status(http.StatusBadRequest).JSON(response)

	}
	c.unitService.Delete(context.Context(), strconv.Itoa(id))
	response := helper.APIResponse("Success to delete data", http.StatusOK, "success", nil)
	return context.Status(http.StatusOK).JSON(response)
}

// FindByID implements UnitController.
func (c *unitController) FindByID(context *fiber.Ctx) error {
	id, err := strconv.Atoi(context.Params("id_unit"))
	if err != nil {
		response := helper.APIResponse("Failed to get data", http.StatusBadRequest, "error", err.Error())
		return context.Status(http.StatusBadRequest).JSON(response)

	}
	unit, err := c.unitService.FindByID(context.Context(), strconv.Itoa(id))

	if err != nil {
		response := helper.APIResponse("Failed to get data", http.StatusBadRequest, "error", err.Error())
		return context.Status(http.StatusBadRequest).JSON(response)
	}

	response := helper.APIResponse("Success to get data", http.StatusOK, "success", unit)
	return context.Status(http.StatusOK).JSON(response)
}

// Update implements UnitController.
func (c *unitController) Update(context *fiber.Ctx) error {
	id, err := strconv.Atoi(context.Params("id_unit"))
	if err != nil {
		response := helper.APIResponse("Failed to put data", http.StatusBadRequest, "error", err.Error())
		return context.Status(http.StatusBadRequest).JSON(response)
	}
	var inputData request.UnitUpdate

	if err := context.BodyParser(&inputData); err != nil {
		response := helper.APIResponse("Failed to put data", http.StatusBadRequest, "error", err.Error())
		return context.Status(http.StatusBadRequest).JSON(response)
	}
	inputData.IdUnit = strconv.Itoa(id)

	errors := request.ValidateStruct(&inputData)
	if errors != nil {
		response := helper.APIResponse("Failed to put data", http.StatusBadRequest, "error", errors)
		return context.Status(http.StatusBadRequest).JSON(response)
	}

	if err := c.unitService.Update(context.Context(), inputData); err != nil {
		response := helper.APIResponse("Failed to put data", http.StatusBadRequest, "error", err.Error())
		return context.Status(http.StatusBadRequest).JSON(response)
	}
	response := helper.APIResponse("Success to put data", http.StatusOK, "success", nil)
	return context.Status(http.StatusOK).JSON(response)
}

// Upload implements UnitController.
func (*unitController) Upload(context *fiber.Ctx) error {
	panic("unimplemented")
}

func NewUnitController(unitService service.UnitService) UnitController {
	return &unitController{
		unitService: unitService,
	}
}
