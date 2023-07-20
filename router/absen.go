package router

import (
	"database/sql"
	"materi/controller"
	"materi/repository"
	"materi/service"

	"github.com/gofiber/fiber/v2"
)

func AbsenRouter(db *sql.DB, base *fiber.App) {
	var (
		AbsenRepository repository.AbsenRepository = repository.NewAbsenRepository(db)
		AbsenService    service.AbsenService       = service.NewAbsenService(AbsenRepository)
		AbsenController controller.AbsenController = controller.NewAbsenController(AbsenService)
	)

	root := base.Group("/Absen")
	root.Get("/", AbsenController.All)
	root.Get("/:id_absen", AbsenController.FindByID)
	root.Post("/", AbsenController.Create)
	root.Put("/:id_absen", AbsenController.Update)
	root.Put("/profil/:id_absen", AbsenController.Upload)
	root.Delete("/:id_absen", AbsenController.Delete)
}
