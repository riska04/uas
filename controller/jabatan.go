package controller

import (
	"materi/helper"
	"materi/service"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type JabatanController interface {
	All(context *fiber.Ctx) error
	FindByID(context *fiber.Ctx) error
	Create(context *fiber.Ctx) error
	Update(context *fiber.Ctx) error
	Delete(context *fiber.Ctx) error
	Upload(context *fiber.Ctx) error
}

type jabatanController struct {
	jabatanService service.JabatanService
}

// All implements JabatanController.
func (c *jabatanController) All(context *fiber.Ctx) error {
	user, err := c.jabatanService.All(context.Context())

	if err != nil {
		response := helper.APIResponse("Failed to get data", http.StatusBadRequest, "error", err.Error())
		return context.Status(http.StatusBadRequest).JSON(response)
	}
	response := helper.APIResponse("Success to get data", http.StatusOK, "success", user)
	return context.Status(http.StatusOK).JSON(response)
}

// Create implements JabatanController.
func (c *jabatanController) Create(context *fiber.Ctx) error {
	panic("unimplemented")
}

// Delete implements JabatanController.
func (c *jabatanController) Delete(context *fiber.Ctx) error {
	panic("unimplemented")
}

// FindByID implements JabatanController.
func (c *jabatanController) FindByID(context *fiber.Ctx) error {
	panic("unimplemented")
}

// Update implements JabatanController.
func (c *jabatanController) Update(context *fiber.Ctx) error {
	panic("unimplemented")
}

// Upload implements JabatanController.
func (c *jabatanController) Upload(context *fiber.Ctx) error {
	panic("unimplemented")
}

func NewJabatanController(jabatanService service.JabatanService) JabatanController {
	return &jabatanController{
		jabatanService: jabatanService,
	}
}
