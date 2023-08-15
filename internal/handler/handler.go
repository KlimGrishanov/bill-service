package handler

import (
	"github.com/KlimGrishanov/bill-service/internal/usecase"
	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	services *usecase.UseCase
}

func NewHandler(services *usecase.UseCase) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoute() *fiber.App {
	router := fiber.New()
	api := router.Group("/bill-api")
	api.Post("/income", h.PostIncomeHandler)
	api.Post("/withdraw", h.PostWithdrawHandler)
	api.Post("/transfer", h.PostTransferHandler)
	api.Get("/balance", h.GetBalanceHandler)

	return router
}
