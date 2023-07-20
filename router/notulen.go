package router

import (
	"database/sql"
	"materi/controller"
	"materi/repository"
	"materi/service"

	"github.com/gofiber/fiber/v2"
)

func NotulenRouter(db *sql.DB, base *fiber.App) {
	var (
		NotulenRepository repository.NotulenRepository = repository.NewNotulenRepository(db)
		NotulenService    service.NotulenService       = service.NewNotulenService(NotulenRepository)
		NotulenController controller.NotulenController = controller.NewNotulenController(NotulenService)
	)

	root := base.Group("/Notulen")
	root.Get("/", NotulenController.All)
	root.Get("/:id_notulen", NotulenController.FindByID)
	root.Post("/", NotulenController.Create)
	root.Put("/:id_notulen", NotulenController.Update)
	root.Put("/profil/:id_notulen", NotulenController.Upload)
	root.Delete("/:id_notulen", NotulenController.Delete)
}
