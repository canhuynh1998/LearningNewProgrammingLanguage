package controllers
import (
	"go-practice/backendApi/services"
	"github.com/gofiber/fiber"
)

func GetBookFromId(c *fiber.Ctx) {
	id := c.Params("id")
	book := services.GetBookFromId(id)
	if book == nil {
		
	}
}