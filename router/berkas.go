package router

import (
	"database/sql"
	"materi/controller"
	"materi/repository"
	"materi/service"

	"github.com/gofiber/fiber/v2"
)

func BerkasRouter(db *sql.DB, base *fiber.App) {
	var (
		BerkasRepository repository.BerkasRepository = repository.NewBerkasRepository(db)
		BerkasService    service.BerkasService       = service.NewBerkasService(BerkasRepository)
		BerkasController controller.BerkasController = controller.NewBerkasController(BerkasService)
	)

	root := base.Group("/Berkas")
	root.Get("/", BerkasController.All)
	root.Get("/:id_berkas", BerkasController.FindByID)
	root.Post("/", BerkasController.Create)
	root.Put("/:id_berkas", BerkasController.Update)
	root.Put("/profil/:id_berkas", BerkasController.Upload)
	root.Delete("/:id_berkas", BerkasController.Delete)
}
