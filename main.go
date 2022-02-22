package main

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/hasimy-as/go-fiber-api/bin"
	"github.com/hasimy-as/go-fiber-api/bin/config"
	"github.com/hasimy-as/go-fiber-api/bin/db"
	"github.com/hasimy-as/go-fiber-api/utils/res"
)

func main() {
	app := fiber.New()
	bin.Route(app)
	db.ConnectDB()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(res.Response{Success: true, Data: "OK", Message: "Server is working properly", Code: http.StatusOK})
	})

	app.Listen(config.GlobalEnv.PORT)
}
