package controller

import (
	"materi/helper"
	"materi/model/request"
	"materi/service"
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type UndanganController interface {
	All(context *fiber.Ctx) error
	FindByID(context *fiber.Ctx) error
	Create(context *fiber.Ctx) error
	Update(context *fiber.Ctx) error
	Delete(context *fiber.Ctx) error
	Upload(context *fiber.Ctx) error
}

type undanganController struct {
	undanganService service.UndanganService
}

// All implements UndanganController.
func (c *undanganController) All(context *fiber.Ctx) error {
	undangan, err := c.undanganService.All(context.Context())

	if err != nil {
		response := helper.APIResponse("Failed to get data", http.StatusBadRequest, "error", err.Error())
		return context.Status(http.StatusBadRequest).JSON(response)
	}
	response := helper.APIResponse("Success to get data", http.StatusOK, "success", undangan)
	return context.Status(http.StatusOK).JSON(response)
}

// Create implements UndanganController.
func (c *undanganController) Create(context *fiber.Ctx) error {
	var inputData request.UndanganCreate
	if err := context.BodyParser(&inputData); err != nil {
		response := helper.APIResponse("Failed to post data", http.StatusBadRequest, "error", err.Error())
		return context.Status(http.StatusBadRequest).JSON(response)
	}
	errors := request.ValidateStruct(&inputData)
	if errors != nil {
		response := helper.APIResponse("Failed to post data", http.StatusBadRequest, "error", errors)
		return context.Status(http.StatusBadRequest).JSON(response)
	}

	_, err := c.undanganService.Create(context.Context(), inputData)
	if err != nil {
		response := helper.APIResponse("Failed to post data", http.StatusBadRequest, "error", err.Error())
		return context.Status(http.StatusBadRequest).JSON(response)
	}
	response := helper.APIResponse("Success to post data", http.StatusOK, "success", nil)
	return context.Status(http.StatusOK).JSON(response)
}

// Delete implements UndanganController.
func (c *undanganController) Delete(context *fiber.Ctx) error {
	id, err := strconv.Atoi(context.Params("id_undangan"))
	if err != nil {
		response := helper.APIResponse("Failed to get data", http.StatusBadRequest, "error", err.Error())
		return context.Status(http.StatusBadRequest).JSON(response)

	}
	c.undanganService.Delete(context.Context(), strconv.Itoa(id))
	response := helper.APIResponse("Success to delete data", http.StatusOK, "success", nil)
	return context.Status(http.StatusOK).JSON(response)
}

// FindByID implements UndanganController.
func (c *undanganController) FindByID(context *fiber.Ctx) error {
	id, err := strconv.Atoi(context.Params("id_undangan"))
	if err != nil {
		response := helper.APIResponse("Failed to get data", http.StatusBadRequest, "error", err.Error())
		return context.Status(http.StatusBadRequest).JSON(response)

	}
	undangan, err := c.undanganService.FindByID(context.Context(), strconv.Itoa(id))

	if err != nil {
		response := helper.APIResponse("Failed to get data", http.StatusBadRequest, "error", err.Error())
		return context.Status(http.StatusBadRequest).JSON(response)
	}

	response := helper.APIResponse("Success to get data", http.StatusOK, "success", undangan)
	return context.Status(http.StatusOK).JSON(response)
}

// Update implements UndanganController.
func (c *undanganController) Update(context *fiber.Ctx) error {
	id, err := strconv.Atoi(context.Params("id_undangan"))
	if err != nil {
		response := helper.APIResponse("Failed to put data", http.StatusBadRequest, "error", err.Error())
		return context.Status(http.StatusBadRequest).JSON(response)
	}
	var inputData request.UndanganUpdate

	if err := context.BodyParser(&inputData); err != nil {
		response := helper.APIResponse("Failed to put data", http.StatusBadRequest, "error", err.Error())
		return context.Status(http.StatusBadRequest).JSON(response)
	}
	inputData.IdUndangan = strconv.Itoa(id)

	errors := request.ValidateStruct(&inputData)
	if errors != nil {
		response := helper.APIResponse("Failed to put data", http.StatusBadRequest, "error", errors)
		return context.Status(http.StatusBadRequest).JSON(response)
	}

	if err := c.undanganService.Update(context.Context(), inputData); err != nil {
		response := helper.APIResponse("Failed to put data", http.StatusBadRequest, "error", err.Error())
		return context.Status(http.StatusBadRequest).JSON(response)
	}
	response := helper.APIResponse("Success to put data", http.StatusOK, "success", nil)
	return context.Status(http.StatusOK).JSON(response)
}

// Upload implements UndanganController.
func (*undanganController) Upload(context *fiber.Ctx) error {
	panic("unimplemented")
}

func NewUndanganController(undanganService service.UndanganService) UndanganController {
	return &undanganController{
		undanganService: undanganService,
	}
}
