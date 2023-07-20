package controller

import (
	"materi/helper"
	"materi/model/request"
	"materi/service"
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type AbsenController interface {
	All(context *fiber.Ctx) error
	FindByID(context *fiber.Ctx) error
	Create(context *fiber.Ctx) error
	Update(context *fiber.Ctx) error
	Delete(context *fiber.Ctx) error
	Upload(context *fiber.Ctx) error
}

type absenController struct {
	absenService service.AbsenService
}

// All implements AbsenController.
func (c *absenController) All(context *fiber.Ctx) error {
	absen, err := c.absenService.All(context.Context())

	if err != nil {
		response := helper.APIResponse("Failed to get data", http.StatusBadRequest, "error", err.Error())
		return context.Status(http.StatusBadRequest).JSON(response)
	}
	response := helper.APIResponse("Success to get data", http.StatusOK, "success", absen)
	return context.Status(http.StatusOK).JSON(response)

}

// Create implements AbsenController.
func (c *absenController) Create(context *fiber.Ctx) error {
	var inputData request.AbsenCreate
	if err := context.BodyParser(&inputData); err != nil {
		response := helper.APIResponse("Failed to post data", http.StatusBadRequest, "error", err.Error())
		return context.Status(http.StatusBadRequest).JSON(response)
	}
	errors := request.ValidateStruct(&inputData)
	if errors != nil {
		response := helper.APIResponse("Failed to post data", http.StatusBadRequest, "error", errors)
		return context.Status(http.StatusBadRequest).JSON(response)
	}

	_, err := c.absenService.Create(context.Context(), inputData)
	if err != nil {
		response := helper.APIResponse("Failed to post data", http.StatusBadRequest, "error", err.Error())
		return context.Status(http.StatusBadRequest).JSON(response)
	}
	response := helper.APIResponse("Success to post data", http.StatusOK, "success", nil)
	return context.Status(http.StatusOK).JSON(response)
}

// Delete implements AbsenController.
func (c *absenController) Delete(context *fiber.Ctx) error {
	id, err := strconv.Atoi(context.Params("id_absen"))
	if err != nil {
		response := helper.APIResponse("Failed to get data", http.StatusBadRequest, "error", err.Error())
		return context.Status(http.StatusBadRequest).JSON(response)

	}
	c.absenService.Delete(context.Context(), strconv.Itoa(id))
	response := helper.APIResponse("Success to delete data", http.StatusOK, "success", nil)
	return context.Status(http.StatusOK).JSON(response)
}

// FindByID implements AbsenController.
func (c *absenController) FindByID(context *fiber.Ctx) error {
	id, err := strconv.Atoi(context.Params("id_absen"))
	if err != nil {
		response := helper.APIResponse("Failed to get data", http.StatusBadRequest, "error", err.Error())
		return context.Status(http.StatusBadRequest).JSON(response)

	}
	absen, err := c.absenService.FindByID(context.Context(), strconv.Itoa(id))

	if err != nil {
		response := helper.APIResponse("Failed to get data", http.StatusBadRequest, "error", err.Error())
		return context.Status(http.StatusBadRequest).JSON(response)
	}

	response := helper.APIResponse("Success to get data", http.StatusOK, "success", absen)
	return context.Status(http.StatusOK).JSON(response)
}

// Update implements AbsenController.
func (c *absenController) Update(context *fiber.Ctx) error {
	id, err := strconv.Atoi(context.Params("id_absen"))
	if err != nil {
		response := helper.APIResponse("Failed to put data", http.StatusBadRequest, "error", err.Error())
		return context.Status(http.StatusBadRequest).JSON(response)
	}
	var inputData request.AbsenUpdate

	if err := context.BodyParser(&inputData); err != nil {
		response := helper.APIResponse("Failed to put data", http.StatusBadRequest, "error", err.Error())
		return context.Status(http.StatusBadRequest).JSON(response)
	}
	inputData.IdAbsen = strconv.Itoa(id)

	errors := request.ValidateStruct(&inputData)
	if errors != nil {
		response := helper.APIResponse("Failed to put data", http.StatusBadRequest, "error", errors)
		return context.Status(http.StatusBadRequest).JSON(response)
	}

	if err := c.absenService.Update(context.Context(), inputData); err != nil {
		response := helper.APIResponse("Failed to put data", http.StatusBadRequest, "error", err.Error())
		return context.Status(http.StatusBadRequest).JSON(response)
	}
	response := helper.APIResponse("Success to put data", http.StatusOK, "success", nil)
	return context.Status(http.StatusOK).JSON(response)
}

// Upload implements BerkasController.
func (c *absenController) Upload(context *fiber.Ctx) error {
	panic("unimplemented")
}

func NewAbsenController(absenService service.AbsenService) AbsenController {
	return &absenController{
		absenService: absenService,
	}
}
