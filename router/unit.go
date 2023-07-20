package router

import (
	"database/sql"
	"materi/controller"
	"materi/repository"
	"materi/service"

	"github.com/gofiber/fiber/v2"
)

func UnitRouter(db *sql.DB, base *fiber.App) {
	var (
		UnitRepository repository.UnitRepository = repository.NewUnitRepository(db)
		UnitService    service.UnitService       = service.NewUnitService(UnitRepository)
		UnitController controller.UnitController = controller.NewUnitController(UnitService)
	)

	root := base.Group("/Unit")
	root.Get("/", UnitController.All)
	root.Get("/:id_unit", UnitController.FindByID)
	root.Post("/", UnitController.Create)
	root.Put("/:id_unit", UnitController.Update)
	root.Put("/profil/:id_unit", UnitController.Upload)
	root.Delete("/:id_unit", UnitController.Delete)
}
