package db

import (
	"fmt"
	"simple-go-search/config"
	"simple-go-search/server/models"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewGormDB() Database {
	return &DB{}
}

func (d *DB) GetConnection() *DB {
	dsn := fmt.Sprintf(`
		host=%s port=%s user=%s password=%s dbname=%s sslmode=%s`,
		config.GetString(config.POSTGRES_HOST, "localhost"),
		config.GetString(config.POSTGRES_PORT, "8888"),
		config.GetString(config.POSTGRES_USER, "search-user"),
		config.GetString(config.POSTGRES_PASS, "search-user"),
		config.GetString(config.POSTGRES_DB_NAME, "search"),
		config.GetString(config.POSTGRES_SSLMODE, "disable"),
	)

	db, err := gorm.Open(postgres.Open(dsn))
	if err != nil {
		panic(err)
	}

	d.GormDB = db

	db.Debug().AutoMigrate(models.Course{})
	return d
}

func (d *DB) SetConnectionPool() {
	sqlDB, err := d.GormDB.DB()
	if err != nil {
		panic(err)
	}

	maxOpen := config.GetInt8(config.POSTGRES_MAX_OPEN_CONN, 20)
	maxLifetime := config.GetInt8(config.POSTGRES_MAX_LIFETIME, 20)
	maxIdleConn := config.GetInt8(config.POSTGRES_MAX_IDLE_CONN, 20)
	maxIdleTime := config.GetInt8(config.POSTGRES_MAX_IDLE_TIME, 20)

	sqlDB.SetMaxOpenConns(int(maxOpen))
	sqlDB.SetConnMaxLifetime(time.Duration(maxLifetime) * time.Second)
	sqlDB.SetMaxIdleConns(int(maxIdleConn))
	sqlDB.SetConnMaxIdleTime(time.Duration(maxIdleTime) * time.Second)

}
