package routers

import (
	"shopping-site/api/handlers"
	"shopping-site/api/middleware"
	"shopping-site/api/repositories"
	"shopping-site/api/services"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func UserRoute(app *fiber.App, db *gorm.DB) {
	userRepository := repositories.CommenceUserRepository(db)

	userService := services.CommenceUserService(userRepository)

	handler := handlers.UserHandler{UserService: userService}

	user := app.Group("/v1/role/users")
	user.Use(middleware.ValidateJwt)

	user.Get("", handler.GetUsersHandler)

}
