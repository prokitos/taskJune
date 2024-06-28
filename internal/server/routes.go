package server

import (
	"module/internal/services"

	"github.com/gofiber/fiber/v2"
)

// получить.  (exchange_name, pair string) => OrderBook. Asks? Bids?
func RouteGetBook(c *fiber.Ctx) error {

	_, res := services.GetOrderBook(c)

	return res

}

// записать (exchange_name pair string []*DepthOrder) => OrderBook. asks? bids?
func RouteSaveBook(c *fiber.Ctx) error {

	return services.SaveOrderBook(c)

}

// получить. (Client) => []HistoryOrder. видимо можно искать только по имени, или по имени и чемуто ещё...
func RouteGetHistory(c *fiber.Ctx) error {

	_, res := services.GetOrderHistory(c)
	return res

}

// записать (Client, HistoryOrder) => Client, HistoryOrder
func RouteSaveClient(c *fiber.Ctx) error {

	return services.SaveOrder(c)

}
