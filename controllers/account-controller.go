package controllers

import (
	"transaction-service/services"

	"github.com/gofiber/fiber/v2"
)

type AccountController struct {
	Service *services.AccountService
}

func NewAccountController(service *services.AccountService) *AccountController {
	return &AccountController{Service: service}
}

func (c *AccountController) Register(ctx *fiber.Ctx) error {
	var body struct {
		Name  string `json:"nama"`
		NIK   string `json:"nik"`
		Phone string `json:"no_hp"`
	}

	if err := ctx.BodyParser(&body); err != nil {
		return ctx.Status(400).JSON(fiber.Map{"remark": "Payload tidak valid"})
	}

	accNum, err := c.Service.Register(body.Name, body.NIK, body.Phone)
	if err != nil {
		return ctx.Status(400).JSON(fiber.Map{"remark": err.Error()})
	}

	return ctx.JSON(fiber.Map{"no_rekening": accNum})
}

func (c *AccountController) Deposit(ctx *fiber.Ctx) error {
	var body struct {
		AccountNumber string  `json:"no_rekening"`
		Amount        float64 `json:"nominal"`
	}

	if err := ctx.BodyParser(&body); err != nil {
		return ctx.Status(400).JSON(fiber.Map{"remark": "Payload tidak valid"})
	}

	saldo, err := c.Service.Deposit(body.AccountNumber, body.Amount)
	if err != nil {
		return ctx.Status(400).JSON(fiber.Map{"remark": err.Error()})
	}

	return ctx.JSON(fiber.Map{"saldo": saldo})
}

func (c *AccountController) Withdraw(ctx *fiber.Ctx) error {
	var body struct {
		AccountNumber string  `json:"no_rekening"`
		Amount        float64 `json:"nominal"`
	}

	if err := ctx.BodyParser(&body); err != nil {
		return ctx.Status(400).JSON(fiber.Map{"remark": "Payload tidak valid"})
	}

	saldo, err := c.Service.Withdraw(body.AccountNumber, body.Amount)
	if err != nil {
		return ctx.Status(400).JSON(fiber.Map{"remark": err.Error()})
	}

	return ctx.JSON(fiber.Map{"saldo": saldo})
}

func (c *AccountController) GetBalance(ctx *fiber.Ctx) error {
	accNum := ctx.Params("no_rekening")

	saldo, err := c.Service.GetBalance(accNum)
	if err != nil {
		return ctx.Status(400).JSON(fiber.Map{"remark": err.Error()})
	}

	return ctx.JSON(fiber.Map{"saldo": saldo})
}
