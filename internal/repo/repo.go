package repo

import (
	"Uvarenko2022/restaurant/database"
	"Uvarenko2022/restaurant/internal/entity"
	"context"
)

type IFoodRepo interface {
	GetProperFood(ids []uint) ([]entity.Food, error)
	GetAllFood() ([]entity.Food, error)
	Put(f *entity.Food) error
	UpdateFood(food *entity.Food) error
}

type IOrderRepo interface {
	CreateOrder(order *entity.Order) error
	GetOrders(limit int, offset int) ([]entity.Order, error)
	UpdateOrder(order *entity.Order) error
}

type ICacheRepo interface {
	UpdateCache(ctx context.Context, food []entity.Food) error
	GetAllCache(ctx context.Context) ([]entity.CacheFood, error)
	GetProperCahce(ctx context.Context, ids []uint) ([]entity.CacheFood, error)
}

type Repository struct {
	Ifoodr    IFoodRepo
	Iordrrepo IOrderRepo
	Ichrepo   ICacheRepo
}

func NewRepository(db *database.Postgre, cache *database.Redis) *Repository {
	return &Repository{
		Ifoodr:    newFoodR(db),
		Iordrrepo: newOrderR(db),
		Ichrepo:   newCacheR(cache),
	}
}
