package router

import (
	"database/sql"
	"materi/controller"
	"materi/repository"
	"materi/service"

	"github.com/gofiber/fiber/v2"
)

func LembagaRouter(db *sql.DB, base *fiber.App) {
	var (
		LembagaRepository repository.LembagaRepository = repository.NewLembagaRepository(db)
		LembagaService    service.LembagaService       = service.NewLembagaService(LembagaRepository)
		LembagaController controller.LembagaController = controller.NewLembagaController(LembagaService)
	)

	root := base.Group("/Lembaga")
	root.Get("/", LembagaController.All)
	root.Get("/:id_lembaga", LembagaController.FindByID)
	root.Post("/", LembagaController.Create)
	root.Put("/:id_lembaga", LembagaController.Update)
	root.Put("/profil/:id_lembaga", LembagaController.Upload)
	root.Delete("/:id_lembaga", LembagaController.Delete)
}
