package router

import (
	"database/sql"
	"materi/controller"
	"materi/repository"
	"materi/service"

	"github.com/gofiber/fiber/v2"
)

func StrukturRouter(db *sql.DB, base *fiber.App) {
	var (
		StrukturRepository repository.StrukturRepository = repository.NewStrukturRepository(db)
		StrukturService    service.StrukturService       = service.NewStrukturService(StrukturRepository)
		StrukturController controller.StrukturController = controller.NewStrukturController(StrukturService)
	)

	root := base.Group("/Struktur")
	root.Get("/", StrukturController.All)
	root.Get("/:id_struktur", StrukturController.FindByID)
	root.Post("/", StrukturController.Create)
	root.Put("/:id_struktur", StrukturController.Update)
	root.Put("/profil/:id_struktur", StrukturController.Upload)
	root.Delete("/:id_struktur", StrukturController.Delete)
}
