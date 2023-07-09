package router

import (
	"database/sql"
	"materi/controller"
	"materi/repository"
	"materi/service"

	"github.com/gofiber/fiber/v2"
)

func JabatanRouter(db *sql.DB, base *fiber.App) {
	var (
		JabatanRepository repository.JabatanRepository = repository.NewJabatanRepository(db)
		JabatanService    service.JabatanService       = service.NewJabatanService(JabatanRepository)
		JabatanController controller.JabatanController = controller.NewJabatanController(JabatanService)
	)

	root := base.Group("/Jabatar")
	root.Get("/", JabatanController.All)
	root.Get("/:id_jabatan", JabatanController.FindByID)
	root.Post("/", JabatanController.Create)
	root.Put("/:id_jabatan", JabatanController.Update)
	root.Put("/profil/:id_jabatan", JabatanController.Upload)
	root.Delete("/:id_jabatan", JabatanController.Delete)
}
