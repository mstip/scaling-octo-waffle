package pkg

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type ExchangeRates struct {
	gorm.Model
	Currency string
	Rate     float32
}

func NewDb() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(&ExchangeRates{})

	if err != nil {
		return nil, err
	}

	return db, nil
}
