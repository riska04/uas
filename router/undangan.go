package router

import (
	"database/sql"
	"materi/controller"
	"materi/repository"
	"materi/service"

	"github.com/gofiber/fiber/v2"
)

func UndanganRouter(db *sql.DB, base *fiber.App) {
	var (
		UndanganRepository repository.UndanganRepository = repository.NewUndanganRepository(db)
		UndanganService    service.UndanganService       = service.NewUndanganService(UndanganRepository)
		UndanganController controller.UndanganController = controller.NewUndanganController(UndanganService)
	)

	root := base.Group("/Undangan")
	root.Get("/", UndanganController.All)
	root.Get("/:id_undangan", UndanganController.FindByID)
	root.Post("/", UndanganController.Create)
	root.Put("/:id_undangan", UndanganController.Update)
	root.Put("/profil/:id_undangan", UndanganController.Upload)
	root.Delete("/:id_undangan", UndanganController.Delete)
}
