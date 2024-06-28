package main

import (
	"module/internal/app"
	"module/internal/config"
	"module/internal/database"
	"os"
	"os/signal"
	"syscall"

	log "github.com/sirupsen/logrus"
)

// @title User API
// @version 1.0
// @description This is a sample service for managing users
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email soberkoder@gmail.com
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8888
// @BasePath /
// @securityDefinitions.apikey Bearer
// @in header
// @name Authorization
// @description Type "Bearer" followed by a space and JWT token.
func main() {

	// установка логов. установка чтобы показывать логи debug уровня
	log.SetLevel(log.DebugLevel)
	log.Info("the server is starting")

	// получение конфигов
	cfg := config.ConfigMustLoad()

	// миграция и подключение к бд.
	err := database.CheckDatabaseCreated(cfg.Connect)
	if err != nil {
		return
	}

	database.OpenConnection(cfg.Connect)
	database.StartMigration()

	// запуск сервера в горутине, чтобы потом нормально звершать приложение
	var application app.App
	go application.NewServer(cfg.Server)

	// в итоге мы обрабатываем завершение приложения, и если мы закрыаем его как либо, то оно выполняет действие из метода Stop
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)
	<-stop
	application.Stop() // безопасное выключение сервера

}
