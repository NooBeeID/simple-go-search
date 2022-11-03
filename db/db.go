package db

import (
	"gorm.io/gorm"
)

type DB struct {
	GormDB *gorm.DB
}

type Database interface {
	GetConnection() *DB
	SetConnectionPool()
}
