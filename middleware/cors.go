package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

// Or extend your config for customization
func Cors() fiber.Handler {
	return cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "GET,POST,HEAD,PUT,DELETE,PATCH",
		AllowHeaders: "Accept, Content-Type, Content-Length, Accept-Encoding, Authorization, ResponseType",
	})
}
