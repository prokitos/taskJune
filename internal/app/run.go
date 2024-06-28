package app

import (
	"fmt"
	"module/internal/models"
	"module/internal/server"
	"time"

	"github.com/gofiber/fiber/v2"
)

type App struct {
	Server *fiber.App
}

// запуск сервера с портом, указаным в конфиге
func (a *App) NewServer(port models.ServerConfig) {
	a.Server = server.ServerStart(port)
}

// безопасное выключение сервера.
func (a *App) Stop() {

	// пытается выключить сервер, а если не получится, то через 60 секунд экстренно сбросит соединение
	fmt.Println("Gracefully shutting down...")
	a.Server.ShutdownWithTimeout(60 * time.Second)

}
