package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func GetDB(dbAdress string) *gorm.DB {
	db, err := gorm.Open(postgres.Open(dbAdress), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}

	seedDb(db)
	return db
}
