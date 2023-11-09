package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/yogawahyudi7/jojonomic/config"
)

func main() {

	app := fiber.New()

	app.Use(logger.New(config.LoggerText()), logger.New(config.LoggerTermial()))

	app.Listen(":3000")
}
