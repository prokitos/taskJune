package services

import (
	"module/internal/database"
	"module/internal/models"

	"github.com/gofiber/fiber/v2"
)

func GetOrderBook(c *fiber.Ctx) ([]*models.DepthOrder, error) {

	exchangeName := c.Query("exchange_name", "")
	pair := c.Query("pair", "")

	_, res := database.ShowBooks(c, exchangeName, pair)
	return nil, res
}

func SaveOrderBook(c *fiber.Ctx) error {

	exchangeName := c.Query("exchange_name", "")
	pair := c.Query("pair", "")

	var orderBooks models.ArrayDepthOrder
	if err := c.BodyParser(&orderBooks); err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	res := database.CreateBooks(c, exchangeName, pair, orderBooks.Array)
	return res
}

func GetOrderHistory(c *fiber.Ctx) ([]*models.HistoryOrder, error) {

	var client models.Client
	if err := c.BodyParser(&client); err != nil {
		return nil, c.SendStatus(fiber.StatusBadRequest)
	}

	_, res := database.ShowHistory(c, client)
	return nil, res
}

func SaveOrder(c *fiber.Ctx) error {

	var clientHistory models.ClientAndHistory
	if err := c.BodyParser(&clientHistory); err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	res := database.CreateOrder(c, clientHistory.Client, clientHistory.History)
	return res
}
