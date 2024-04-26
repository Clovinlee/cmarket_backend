package config

import (
	"chris/cmarket/interfaces"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type PostgreDB struct {
	interfaces.IDatabaseInitialize
	db *gorm.DB
}

func NewPostgreDB() *PostgreDB {
	return &PostgreDB{}
}

func (provider *PostgreDB) GetDB() *gorm.DB {
	if provider.db == nil {
		provider.initializeDB()
	}
	return provider.db
}

func (postgreDB *PostgreDB) initializeDB() {
	var err error
	dsn := os.Getenv("DATABASE_URL")

	var db *gorm.DB
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info), //enable info debug log sql
	})

	if err != nil {
		log.Fatal("Error connecting to database. Error: ", err)
	}

	postgreDB.db = db
}
