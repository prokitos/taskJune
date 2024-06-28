package server

import (
	"module/internal/services"

	"github.com/gofiber/fiber/v2"
)

// Users godoc
// @Summary get DepthOrder by name or pair
// @Description get DepthOrder by name or pair
// @Tags DepthOrder
// @Accept  json
// @Produce  json
// @Param exchange_name query string false "Show by exchange_name"
// @Param pair query string false "Show by pair"
// @Success 200 "successful operation"
// @Router /GetOrderBook [get]
func RouteGetBook(c *fiber.Ctx) error {

	res, err := services.GetOrderBook(c)
	return c.Status(fiber.StatusAccepted).JSON(fiber.Map{"status": err.Error(), "data": res})
}

// Users godoc
// @Summary insert orderBook
// @Description insert orderBook
// @Tags DepthOrder
// @Accept  json
// @Produce  json
// @Param exchange_name query string false "insert order"
// @Param pair query string false "insert order"
// @Param orderBook body models.ArrayDepthOrderSwag true "Insert order"
// @Success 200 "successful operation"
// @Router /SaveOrderBook [post]
func RouteSaveBook(c *fiber.Ctx) error {

	err := services.SaveOrderBook(c)
	return c.Status(fiber.StatusAccepted).JSON(fiber.Map{"status": err.Error()})
}

// Users godoc
// @Summary get order History
// @Description get order History
// @Tags DepthOrder
// @Accept  json
// @Produce  json
// @Param client query models.ClientSwag false "Show by client"
// @Success 200 "successful operation"
// @Router /GetOrderHistory [get]
func RouteGetHistory(c *fiber.Ctx) error {

	res, err := services.GetOrderHistory(c)
	return c.Status(fiber.StatusAccepted).JSON(fiber.Map{"status": err.Error(), "data": res})
}

// Users godoc
// @Summary save order
// @Description save order
// @Tags DepthOrder
// @Accept  json
// @Produce  json
// @Param clientHistory body models.ClientAndHistorySwag false "insert client and history"
// @Success 200 "successful operation"
// @Router /SaveOrder [post]
func RouteSaveClient(c *fiber.Ctx) error {

	err := services.SaveOrder(c)
	return c.Status(fiber.StatusAccepted).JSON(fiber.Map{"status": err.Error()})
}
