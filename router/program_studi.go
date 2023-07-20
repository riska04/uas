package router

import (
	"database/sql"
	"materi/controller"
	"materi/repository"
	"materi/service"

	"github.com/gofiber/fiber/v2"
)

func ProgramStudiRouter(db *sql.DB, base *fiber.App) {
	var (
		ProgramStudiRepository repository.ProgramStudiRepository = repository.NewProgramStudiRepository(db)
		ProgramStudiService    service.ProgramStudiService       = service.NewProgramStudiService(ProgramStudiRepository)
		ProgramStudiController controller.ProgramStudiController = controller.NewProgramStudiController(ProgramStudiService)
	)

	root := base.Group("/ProgramStudi")
	root.Get("/", ProgramStudiController.All)
	root.Get("/:id_programstudi", ProgramStudiController.FindByID)
	root.Post("/", ProgramStudiController.Create)
	root.Put("/:id_programstudi", ProgramStudiController.Update)
	root.Put("/profil/:id_programstudi", ProgramStudiController.Upload)
	root.Delete("/:id_programstudi", ProgramStudiController.Delete)
}
