package routers

import (
	"shopping-site/api/handlers"
	"shopping-site/api/middleware"
	"shopping-site/api/repositories"
	"shopping-site/api/services"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func MerchantRoute(app *fiber.App, db *gorm.DB) {
	merchantRepository := repositories.CommenceMerchantRepository(db)

	merchantService := services.CommenceMerchantService(merchantRepository)

	handler := handlers.MerchantHandler{MerchantService: merchantService}

	user := app.Group("/v1/role/merchant")
	user.Use(middleware.ValidateJwt, middleware.MerchantRoleAuthentication)

	user.Post("/product", handler.AddProductHandler)
	user.Delete("/product/:id", handler.RemoveProductHandler)
	user.Patch("/product", handler.UpdateProductHandler)
	user.Patch("", handler.UpdateMerchantHandler)
	user.Patch("order:id", handler.UpdateOrderStatusHandler)
	// user.Get("order", handler.GetOrderHandler)
	user.Get("product", handler.GetProductsHandler)
	user.Get("product:id", handler.GetProductHandler)
}
