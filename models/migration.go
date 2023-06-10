package models

import (
	"log"

	"github.com/shurlz/bank-backend/utils"
)

func InitialMigrations(models ...interface{}) {
	db, err := utils.FetchDatabase()
	if err != nil {
		log.Fatal(err)
	}
	for _, model := range models {
		db.AutoMigrate(model)
	}
}
