package handler

import (
	"github.com/gofiber/fiber/v2"
)

func (h *Handler) PostIncomeHandler(c *fiber.Ctx) error {
	return c.SendString("I'm a POST request!")
}

func (h *Handler) PostWithdrawHandler(c *fiber.Ctx) error {
	return c.SendString("I'm a POST request!")
}

func (h *Handler) PostTransferHandler(c *fiber.Ctx) error {
	return c.SendString("I'm a POST request!")
}
