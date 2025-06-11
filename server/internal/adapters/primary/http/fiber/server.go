package fiber

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func StartFiberServer() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("🚀 Servidor levantado exitosamente")
	})

	port := ":3000"
	fmt.Println("🚀 Servidor Fiber corriendo en http://localhost" + port)
	if err := app.Listen(port); err != nil {
		panic(err)
	}
}
