package routes

import (
	"github.com/gofiber/fiber/v3"

	"github.com/t-saturn/central-user-manager/internal/adapters/handlers"
	"github.com/t-saturn/central-user-manager/internal/adapters/repositories/postgres"
	"github.com/t-saturn/central-user-manager/internal/core/services"
)

func ApplicationRoutes(api fiber.Router) {
	/** ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------- */
	application_repository := postgres.NewApplicationRepository()
	application_service := services.NewApplicationService(application_repository)
	application_handler := handlers.NewApplicationHandler(application_service)

	group := api.Group("/applications")
	group.Post("/", application_handler.Create())
	group.Get("/:id", application_handler.GetByID())
	group.Patch("/:id", application_handler.Update())
	/** ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------- */
	application_role_repository := postgres.NewApplicationRoleRepository()
	application_role_service := services.NewApplicationRoleService(application_role_repository)
	application_role_handler := handlers.NewApplicationRoleHanlder(application_role_service)

	group = api.Group("/application-roles")
	group.Post("/", application_role_handler.Create())
	// group.Get("/:id", application_role_handler.GetByID())
	/** ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------- */

}
