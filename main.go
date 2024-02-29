package main

import (
	"flag"
	"github.com/Ledja22/hotel-reservation/api"
	"github.com/gofiber/fiber/v2"
)

func main() {
	listenAddr := flag.String("listenAddr", ":5000", "The listen address of the API server")
	app := fiber.New()
	appv1 := app.Group("/api/v1")
	appv1.Get("/user", api.HandleGetUser)
	appv1.Get("/users/", api.HandleGetUsers)
	app.Listen(*listenAddr)
}
