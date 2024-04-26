package interfaces

import "gorm.io/gorm"

type IDatabaseInitialize interface {
	InitializeDB()
	GetDB() *gorm.DB
}
