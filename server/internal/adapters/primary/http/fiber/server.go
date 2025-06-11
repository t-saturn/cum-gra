package fiber

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/t-saturn/central-user-manager/server/internal/adapters/primary/http/fiber/middleware"
	"github.com/t-saturn/central-user-manager/server/internal/adapters/primary/http/fiber/routes"
)

func StartFiberServer() {
	app := fiber.New()

	// Middleware de logging (archivo)
	app.Use(middleware.SetupLogger())

	// También mostrar por consola (opcional)
	app.Use(func(c *fiber.Ctx) error {
		fmt.Printf("[Fiber] %s %s\n", c.Method(), c.Path())
		return c.Next()
	})

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Servidor levantado exitosamente")
	})

	routes.SetupRoutes(app)

	port := ":3000"
	fmt.Println("Servidor Fiber corriendo en http://localhost" + port)
	if err := app.Listen(port); err != nil {
		panic(err)
	}
}
