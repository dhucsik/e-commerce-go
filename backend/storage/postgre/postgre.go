package postgre

import (
	"context"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Dial(ctx context.Context, url string) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(url), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	if db != nil {
		db.AutoMigrate(&User{}, &Category{}, &Product{}, &Review{}, &Cart{}, &Order{}, &OrderedItem{}, &Payment{})
	}

	return db, nil
}
