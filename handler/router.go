package handler

import (
	"context"
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
)

var (
	d struct{}
)

type Ihandler interface {
	Init() error
}

type Handler struct {
	router *fiber.App
	ctx    context.Context
}

func NewHandler(router *fiber.App, ctx context.Context) error {

	handler := &Handler{
		router: router,
		ctx:    ctx,
	}

	if err := handler.Init(); err != nil {
		return err
	}

	return nil
}

func (h Handler) Init() error {

	// middleware
	// h.router.Use(middleware.Logger())

	// h.router.Use(middleware.Limiter())

	// h.router.Use(middleware.Cors())

	// h.router.Use(adaptor.HTTPMiddleware(middleware.TracingMiddleware))

	// health check
	h.router.Get("health-check", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{
			"IsSuccess": true,
			"Message":   fmt.Sprintf("SVC RUNNING %s - Version %s", os.Getenv("SERVICE_NAME"), os.Getenv("SERVICE_VERSION")),
			"Data":      d,
			"Status":    "0",
		})
	})
	h.Graphql()

	return nil
}
