package routes

import (
	"server/internal/handlers/roles"

	"github.com/gofiber/fiber/v3"
)

func RegisterRolesAssignmentsRoutes(router fiber.Router) {
	app := router.Group("/roles-assignments")

	app.Get("/", handlers.GetRoleAssignmentsHandler)
	app.Get("/stats", handlers.GetRoleAssignmentsStatsHandler)
}
