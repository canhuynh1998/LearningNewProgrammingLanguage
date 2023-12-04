package routes
import (
	"github.com/gofiber/fiber"
	"go-practice/backendApi/services"
)

// Handler
func HelloRoute(a *fiber.App) {
	route := a.Group("/test")

	route.Get("*", services.TestService)
}
