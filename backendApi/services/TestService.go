package services

import "github.com/gofiber/fiber"

func TestService(c *fiber.Ctx) {
	c.Send("hello")
}