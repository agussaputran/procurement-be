package main

import (
	"context"
	"log"
	cf "procurement-be/config"
	h "procurement-be/handler"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

var (
	conf = cf.Config{}
	d    struct{}
	err  error
)

func run() error {
	if err = godotenv.Load(); err != nil {
		return err
	}

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// start server
	server := fiber.New(
		fiber.Config{
			ErrorHandler: func(c *fiber.Ctx, err error) error {
				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
					"Status":    fiber.StatusInternalServerError,
					"IsSuccess": false,
					"Message":   err.Error(),
					"Data":      d,
				})
			},
		},
	)

	// register handler
	if err = h.NewHandler(server, ctx); err != nil {
		return err
	}

	port := "8080"
	if conf.Service.Port != 0 {
		port = strconv.Itoa(conf.Service.Port)
	}
	return server.Listen("0.0.0.0:" + port)
}

func main() {
	if err = run(); err != nil {
		log.Fatal(err)
	}
}
