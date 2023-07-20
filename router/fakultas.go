package router

import (
	"database/sql"
	"materi/controller"
	"materi/repository"
	"materi/service"

	"github.com/gofiber/fiber/v2"
)

func FakultasRouter(db *sql.DB, base *fiber.App) {
	var (
		FakultasRepository repository.FakultasRepository = repository.NewFakultasRepository(db)
		FakultasService    service.FakultasService       = service.NewFakultasService(FakultasRepository)
		FakultasController controller.FakultasController = controller.NewFakultasController(FakultasService)
	)

	root := base.Group("/Fakultas")
	root.Get("/", FakultasController.All)
	root.Get("/:id_fakultas", FakultasController.FindByID)
	root.Post("/", FakultasController.Create)
	root.Put("/:id_fakultas", FakultasController.Update)
	root.Put("/profil/:id_fakultas", FakultasController.Upload)
	root.Delete("/:id_fakultas", FakultasController.Delete)

}
