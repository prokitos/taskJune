package database

import (
	"module/internal/models"

	"github.com/gofiber/fiber/v2"
	log "github.com/sirupsen/logrus"
)

// показать цены покупки/продажи
func ShowBooks(c *fiber.Ctx, exchangeName string, pair string) ([]models.DepthOrder, error) {

	var finded []models.OrderBook
	var searcher models.OrderBook
	searcher.Exchange = exchangeName
	searcher.Pair = pair

	// выполнить поиск
	results := GlobalHandler.Preload("Asks").Find(&finded, searcher)
	if results.Error != nil {
		log.Error("Error find at ShowBooks")
		log.Debug(results.Error)
		return nil, models.ResponseErrorAtServer()
	}

	// проверка что что-то нашлось
	if len(finded) == 0 {
		log.Debug("nothing to show")
		return nil, models.ResponseBadRequest()
	}

	return finded[0].Asks, models.ResponseGood()
}

// записать цены покупки/продажи
func CreateBooks(c *fiber.Ctx, exchange_name string, pair string, records []models.DepthOrder) error {

	// возможно ещё както отсортировать records на asks и bids

	var inserted models.OrderBook
	inserted.Exchange = exchange_name
	inserted.Pair = pair
	inserted.Asks = records
	inserted.Bids = records

	// выпонение вставки
	if result := GlobalHandler.Create(&inserted); result.Error != nil {
		log.Error("Error insert at CreateBooks")
		log.Debug(result.Error)
		return models.ResponseErrorAtServer()
	}

	return models.ResponseGood()
}

// показать историю транзакций
func ShowHistory(c *fiber.Ctx, client models.Client) ([]models.HistoryOrder, error) {

	var finded []models.HistoryOrder
	var searcher models.HistoryOrder
	searcher.Client_name = client.Client_name
	searcher.Exchange_name = client.Exchange_name
	searcher.Label = client.Label
	searcher.Pair = client.Pair

	// выполнить поиск
	results := GlobalHandler.Find(&finded, searcher)
	if results.Error != nil {
		log.Error("Error find at ShowHistory")
		log.Debug(results.Error)
		return nil, models.ResponseErrorAtServer()
	}

	// проверка что что-то нашлось
	if len(finded) == 0 {
		log.Debug("nothing to show")
		return nil, models.ResponseBadRequest()
	}

	return finded, models.ResponseGood()
}

// добавить ордер в историю и клиенту
func CreateOrder(c *fiber.Ctx, client models.Client, history models.HistoryOrder) error {

	// добавление в таблицу клиентов
	if result := GlobalHandler.Create(&client); result.Error != nil {
		log.Error("Error insert client at CreateOrder")
		log.Debug(result.Error)
		return models.ResponseErrorAtServer()
	}

	// добавление в таблицу истории
	if result := GlobalHandler.Create(&history); result.Error != nil {
		log.Error("Error insert history at CreateOrder")
		log.Debug(result.Error)
		return models.ResponseErrorAtServer()
	}

	return models.ResponseGood()
}
