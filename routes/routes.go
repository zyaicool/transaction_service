package routes

import (
	"transaction-service/config"
	"transaction-service/controllers"
	"transaction-service/repositories"
	"transaction-service/services"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {

	repo := repositories.NewAccountRepository(config.DB)
	service := services.NewAccountService(repo, config.DB)
	controller := controllers.NewAccountController(service)

	app.Post("/daftar", controller.Register)
	app.Post("/tabung", controller.Deposit)
	app.Post("/tarik", controller.Withdraw)
	app.Get("/saldo/:no_rekening", controller.GetBalance)
}
