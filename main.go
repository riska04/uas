package main

import (
	"log"
	"materi/config"
	"materi/router"

	"github.com/gofiber/fiber/v2"
)

func main() {

	app := fiber.New(fiber.Config{
		Prefork:       true,
		CaseSensitive: true,
		StrictRouting: true,
	})
	db := config.NewDB()
	router.JabatanRouter(db, app)
	router.AbsenRouter(db, app)
	router.BerkasRouter(db, app)
	router.FakultasRouter(db, app)
	router.LembagaRouter(db, app)
	router.NotulenRouter(db, app)
	router.ProgramStudiRouter(db, app)
	router.StrukturRouter(db, app)
	router.UndanganRouter(db, app)
	router.UnitRouter(db, app)
	router.UserRouter(db, app)
	if err := app.Listen(":3000"); err != nil {
		log.Fatalln(err)
	}
}
