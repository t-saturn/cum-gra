// Package routes contiene las definiciones de rutas HTTP de la aplicación.
package routes

import (
	"github.com/gofiber/fiber/v3"
)

// RegisterRoutes configura y registra todas las rutas HTTP de la aplicación.
func RegisterRoutes(app *fiber.App) {
	app.Get("/", func(c fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "auth service API is running",
		})
	})

	RegisterAuthRoutes(app)
	RegisterHealthRoutes(app)
}

// |   Método   | Ruta                   | Descripción                                                                                                               |
// | ---------- | ---------------------- | ------------------------------------------------------------------------------------------------------------------------- |
// | [ ] `POST` | `/auth/verify`         | Verifica `email` y `password` contra PostgreSQL. Si es válido, responde con `user_id`. (Autenticación base)               |
// | [ ] `POST` | `/auth/login`          | Usa `/auth/verify`, crea sesión activa, genera `access_token`, `refresh_token`, guarda en MongoDB y responde al frontend. |
// | [ ] `POST` | `/auth/logout`         | Invalida `access_token` o `session_id`, marcándolo como revocado en MongoDB.                                              |
// | [ ] `POST` | `/auth/token/refresh`  | Valida el `refresh_token`, genera nuevos tokens y actualiza sesión.                                                       |
// | [ ] `GET`  | `/auth/session/me`     | Retorna información del usuario autenticado, basado en token recibido.                                                    |
// | [ ] `POST` | `/auth/token/validate` | Valida si un `access_token` es válido y activo (útil para frontend, gateway o microservicios).                            |
// | [ ] `POST` | `/auth/logs`           | Registra intentos fallidos o correctos de login. Puede incluir geolocalización, IP, navegador.                            |
// | [ ] `GET`  | `/health`              | Endpoint simple para monitoreo o readiness probe (DevOps).                                                                |
