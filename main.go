package main

import (
	"log"
	"meeting3/database"
	"meeting3/middleware"
	"meeting3/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)


func main(){

	database.ConnectDatabase()

	app:= fiber.New()

	app.Use(
		cors.New(),
		logger.New(),
)

	//Kumpulan route
	app.Get("/", func (c *fiber.Ctx) error {

		//-------Logika Bisnis dari aplikasi/rout

        return c.SendString("Hello, World!")
    })

	app.Get("/hello", routes.Hello)
	app.Get("/heelo/:id", routes.HelloParams)
	app.Post("/register", routes.Register)

	app.Get("/allDat", routes.AllData)


	app.Post("/login", routes.Login)

	//menggunakan middlware 
	app.Put("/update", middleware.JwtProtect(), routes.UpdateDataUser)

	log.Fatal(app.Listen(":3000"))

}