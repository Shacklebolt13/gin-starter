package sql

import (
	"gin-starter/internal/singleton/config"
	"log"
	"sync"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var dbConn *gorm.DB
var dbOnce sync.Once

func ConnectDb(appConfig *config.AppConfig) *gorm.DB {
	var err error

	dbOnce.Do(func() {
		dbConn, err = gorm.Open(sqlite.Open(appConfig.Env.SQLITE_PATH), &gorm.Config{})
		if err != nil {
			log.Fatalf("failed to connect to database: %v", err)
		}
	})

	return dbConn
}
