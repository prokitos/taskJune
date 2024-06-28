package services

import (
	"module/internal/database"
	"module/internal/models"

	"github.com/gofiber/fiber/v2"
	log "github.com/sirupsen/logrus"
)

// показать цены покупки/продажи
func GetOrderBook(c *fiber.Ctx) ([]models.DepthOrder, error) {

	exchangeName := c.Query("exchange_name", "")
	pair := c.Query("pair", "")

	return database.ShowBooks(c, exchangeName, pair)
}

// записать цены покупки/продажи
func SaveOrderBook(c *fiber.Ctx) error {

	exchangeName := c.Query("exchange_name", "")
	pair := c.Query("pair", "")

	var orderBooks models.ArrayDepthOrder
	if err := c.BodyParser(&orderBooks); err != nil {
		log.Error("Error body parsing at SaveOrderBook")
		log.Debug(err)
		return models.ResponseBadRequest()
	}

	return database.CreateBooks(c, exchangeName, pair, orderBooks.Array)
}

// показать историю транзакций
func GetOrderHistory(c *fiber.Ctx) ([]models.HistoryOrder, error) {

	var client models.Client
	client.Client_name = c.Query("client_name", "")
	client.Exchange_name = c.Query("exchange_name", "")
	client.Label = c.Query("label", "")
	client.Pair = c.Query("pair", "")

	return database.ShowHistory(c, client)
}

// добавить ордер в историю и клиенту
func SaveOrder(c *fiber.Ctx) error {

	var clientHistory models.ClientAndHistory
	if err := c.BodyParser(&clientHistory); err != nil {
		log.Error("Error body parsing at GetOrderHistory")
		log.Debug(err)
		return models.ResponseBadRequest()
	}

	return database.CreateOrder(c, clientHistory.Client, clientHistory.History)
}
