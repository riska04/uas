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
	if err := app.Listen(":3000"); err != nil {
		log.Fatalln(err)
	}
}
