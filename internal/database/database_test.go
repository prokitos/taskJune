package database

import (
	"module/internal/models"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/valyala/fasthttp"
)

func TestConnetion(t *testing.T) {

	var cfg models.ConnectConfig
	cfg.Host = "localhost"
	cfg.Name = "postgres"
	cfg.Pass = "root"
	cfg.Port = "8092"
	cfg.User = "postgres"
	cfg.Reload = false

	OpenConnection(cfg)
	sqlDB, err := GlobalHandler.DB()

	if err != nil {
		t.Errorf(err.Error())
	}

	qq := sqlDB.Stats()

	if qq.Idle == 0 && qq.InUse == 0 {
		t.Errorf("result wrong at test, does not connect to server")
	}

	sqlDB.Close()

}

func TestCreateBooks(t *testing.T) {

	var cfg models.ConnectConfig
	cfg.Host = "localhost"
	cfg.Name = "postgres"
	cfg.Pass = "root"
	cfg.Port = "8092"
	cfg.User = "postgres"
	cfg.Reload = false
	OpenConnection(cfg)

	GlobalHandler.AutoMigrate(models.HistoryOrder{})
	GlobalHandler.AutoMigrate(models.Client{})
	GlobalHandler.AutoMigrate(models.DepthOrder{})
	GlobalHandler.AutoMigrate(models.OrderBook{})

	app := fiber.New()
	c := app.AcquireCtx(&fasthttp.RequestCtx{})

	err := CreateBooks(c, "test", "test", []models.DepthOrder{})

	if err.Error() != models.ResponseGood().Error() {
		t.Errorf("result wrong at test")
	}

}

func TestCreateOrder(t *testing.T) {

	var cfg models.ConnectConfig
	cfg.Host = "localhost"
	cfg.Name = "postgres"
	cfg.Pass = "root"
	cfg.Port = "8092"
	cfg.User = "postgres"
	cfg.Reload = false
	OpenConnection(cfg)

	GlobalHandler.AutoMigrate(models.HistoryOrder{})
	GlobalHandler.AutoMigrate(models.Client{})
	GlobalHandler.AutoMigrate(models.DepthOrder{})
	GlobalHandler.AutoMigrate(models.OrderBook{})

	app := fiber.New()
	c := app.AcquireCtx(&fasthttp.RequestCtx{})

	err := CreateOrder(c, models.Client{}, models.HistoryOrder{})

	if err.Error() != models.ResponseGood().Error() {
		t.Errorf("result wrong at test")
	}

}

func TestShowBook(t *testing.T) {

	var cfg models.ConnectConfig
	cfg.Host = "localhost"
	cfg.Name = "postgres"
	cfg.Pass = "root"
	cfg.Port = "8092"
	cfg.User = "postgres"
	cfg.Reload = false
	OpenConnection(cfg)

	GlobalHandler.AutoMigrate(models.HistoryOrder{})
	GlobalHandler.AutoMigrate(models.Client{})
	GlobalHandler.AutoMigrate(models.DepthOrder{})
	GlobalHandler.AutoMigrate(models.OrderBook{})

	app := fiber.New()
	c := app.AcquireCtx(&fasthttp.RequestCtx{})

	_, err := ShowBooks(c, "test", "test")

	if err.Error() != models.ResponseGood().Error() {
		t.Errorf("result wrong at test")
	}

}

func TestShowHistory(t *testing.T) {

	var cfg models.ConnectConfig
	cfg.Host = "localhost"
	cfg.Name = "postgres"
	cfg.Pass = "root"
	cfg.Port = "8092"
	cfg.User = "postgres"
	cfg.Reload = false
	OpenConnection(cfg)

	GlobalHandler.AutoMigrate(models.HistoryOrder{})
	GlobalHandler.AutoMigrate(models.Client{})
	GlobalHandler.AutoMigrate(models.DepthOrder{})
	GlobalHandler.AutoMigrate(models.OrderBook{})

	app := fiber.New()
	c := app.AcquireCtx(&fasthttp.RequestCtx{})

	_, err := ShowHistory(c, models.Client{})

	if err.Error() != models.ResponseGood().Error() {
		t.Errorf("result wrong at test")
	}

}

// мало времени на полное тестирование этих функций, поэтому тестирую что идёт нормальное обращение к МОК таблицам, и происходят разные действия.
