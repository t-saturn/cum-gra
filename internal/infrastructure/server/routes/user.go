package routes

import (
	"github.com/gofiber/fiber/v3"

	"github.com/t-saturn/central-user-manager/internal/adapters/handlers"
	"github.com/t-saturn/central-user-manager/internal/adapters/repositories/postgres"
	"github.com/t-saturn/central-user-manager/internal/core/services"
	"github.com/t-saturn/central-user-manager/internal/shared/security"
)

func UserRoutes(api fiber.Router) {
	/** ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------- */
	structural_repository := postgres.NewStructuralPositionRepository()
	structural_service := services.NewStructuralPositionService(structural_repository)
	structural_handler := handlers.NewStructuralPositionHandler(structural_service)

	group := api.Group("/structural-positions")
	group.Post("/", structural_handler.Create())
	group.Get("/:id", structural_handler.GetByID())
	group.Patch("/:id", structural_handler.Update())

	/** ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------- */
	organic_repository := postgres.NewOrganicUnitRepository()
	organic_service := services.NewOrganicUnitService(organic_repository)
	organic_handler := handlers.NewOrganicUnitHandler(organic_service)

	group = api.Group("/organic-units")
	group.Post("/", organic_handler.Create())
	group.Get("/:id", organic_handler.GetByID())
	group.Patch("/:id", organic_handler.Update())

	/** ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------- */
	user_repository := postgres.NewUserRepository()
	hasher := security.NewArgon2Service()
	user_service := services.NewUserService(user_repository, hasher)
	user_handler := handlers.NewUserHandler(user_service)

	group = api.Group("/users")
	group.Post("/", user_handler.Create())
	group.Get("/:id", user_handler.GetByID())
	group.Patch("/:id", user_handler.Update())
}
