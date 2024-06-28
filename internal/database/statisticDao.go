package database

import (
	"module/internal/models"

	"github.com/gofiber/fiber/v2"
)

func ShowBooks(c *fiber.Ctx, exchangeName string, pair string) ([]models.DepthOrder, error) {

	var finded []models.OrderBook

	var searcher models.OrderBook
	searcher.Exchange = exchangeName
	searcher.Pair = pair

	results := GlobalHandler.Preload("Asks").Preload("Bids").Find(&finded, searcher)
	if results.Error != nil {
		return nil, c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "error", "message": results.Error})
	}

	if len(finded) == 0 {
		return nil, c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "error", "message": results.Error})
	}

	return nil, c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "asks": finded[0]})

}

func CreateBooks(c *fiber.Ctx, exchange_name string, pair string, records []models.DepthOrder) error {

	// возможно както отсортировать records на asks и bids

	var inserted models.OrderBook
	inserted.Exchange = exchange_name
	inserted.Pair = pair
	inserted.Asks = records
	inserted.Bids = records

	if result := GlobalHandler.Create(&inserted); result.Error != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "error", "message": result.Error.Error()})
	}

	// Send a 201 created response
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"status": "success"})

}

func ShowHistory(c *fiber.Ctx, client models.Client) ([]models.HistoryOrder, error) {

	var finded []models.HistoryOrder

	var searcher models.HistoryOrder
	searcher.Client_name = client.Client_name
	searcher.Exchange_name = client.Exchange_name
	searcher.Label = client.Label
	searcher.Pair = client.Pair

	results := GlobalHandler.Find(&finded, searcher)
	if results.Error != nil {
		return nil, c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "error", "message": results.Error})
	}

	return nil, c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "notes": finded})
}

func CreateOrder(c *fiber.Ctx, client models.Client, history models.HistoryOrder) error {

	if result := GlobalHandler.Create(&client); result.Error != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "error", "message": result.Error.Error()})
	}

	if result := GlobalHandler.Create(&history); result.Error != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "error", "message": result.Error.Error()})
	}

	// Send a 201 created response
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"status": "success"})

}
