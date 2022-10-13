package database

import (
	"Uvarenko2022/restaurant/internal/entity"
	"context"
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Postgre struct {
	*gorm.DB
}

func NewPostgre() *Postgre {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("POSTGRES_DB_HOST"),
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_DB"),
		os.Getenv("POSTGRES_DB_PORT"),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("cannot connect to", err)
	}

	if err := db.SetupJoinTable(&entity.Food{}, "Ingridients", &entity.FoodIng{}); err != nil {
		log.Println(err)
		panic(err)
	}
	if err := db.AutoMigrate(&entity.Order{}, &entity.Food{}, &entity.Ingridient{}, &entity.FoodIng{}); err != nil {
		panic(err)
	}

	return &Postgre{db}
}

func SetUpHookCache(db *Postgre, ctx context.Context, food func(ids []uint) ([]entity.Food, error), f func(ctx context.Context, food []entity.Food) error) {
	db.Callback().Create().After("gorm:create").Register("setcache", func(gormdb *gorm.DB) {
		fd, err := food([]uint{})

		if err != nil {
			fmt.Println(err)
		}

		if err := f(ctx, fd); err != nil {
			log.Fatal(err)
			return
		}
	})
}
