package database

import (
	"fmt"
	"module/internal/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	log "github.com/sirupsen/logrus"
)

var GlobalHandler *gorm.DB

// проверка если есть база данных vortext. если нет, то создать.
func CheckDatabaseCreated(config models.ConnectConfig) {

	connectStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", config.User, config.Pass, config.Host, config.Port, "postgres")

	db, err := gorm.Open(postgres.Open(connectStr), &gorm.Config{})
	if err != nil {
		fmt.Println("error, dont opened")
		return
	}

	stmt := fmt.Sprintf("SELECT * FROM pg_database WHERE datname = '%s';", config.Name)
	rs := db.Raw(stmt)
	if rs.Error != nil {
		fmt.Println("error, dont read bd")
		return
	}

	var rec = make(map[string]interface{})
	if rs.Find(rec); len(rec) == 0 {
		stmt := fmt.Sprintf("CREATE DATABASE %s;", config.Name)
		if rs := db.Exec(stmt); rs.Error != nil {
			fmt.Println("error, dont create a database")
			return
		}
	}

	sql, err := db.DB()
	defer func() {
		_ = sql.Close()
	}()

}

// открыть соединение, и поместить его в глобальную переменну.
func OpenConnection(config models.ConnectConfig) {
	connectStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", config.User, config.Pass, config.Host, config.Port, config.Name)

	db, err := gorm.Open(postgres.Open(connectStr), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}

	GlobalHandler = db
}

// миграция
func StartMigration() {

	GlobalHandler.AutoMigrate(models.HistoryOrder{})
	GlobalHandler.AutoMigrate(models.Client{})
	//GlobalHandler.AutoMigrate(models.OrderBook{})

	log.Info("migration complete")

}
