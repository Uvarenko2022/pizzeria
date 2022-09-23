package usecase

import (
	"Uvarenko2022/restaurant/internal/entity"
	"Uvarenko2022/restaurant/internal/repo"
	"context"
)

//go:generate mockgen -source=usecase.go -destination=mock/mock.go

type PiizzaUseCase struct {
	IFoodUC
	IOrderUC
	ICacheUC
}

func New(repo *repo.Repository) *PiizzaUseCase {
	return &PiizzaUseCase{
		NewFoodUC(repo.Ifoodr),
		NewOrderUC(repo.Iordrrepo),
		NewCacheUC(repo.Ichrepo),
	}
}

type IFoodUC interface {
	GetFood(ids []uint) ([]entity.Food, error)
	AddFood(food *entity.Food) error
	UpdateFood(food *entity.Food) error
}

type IOrderUC interface {
	CreateOrder(order *entity.Order, food []entity.Food) error
	GetOrders(limit int, offset int) ([]entity.Order, error)
	UpdateOrder(order *entity.Order, food []entity.Food) error
}

type ICacheUC interface {
	UpdateCache(ctx context.Context, food []entity.Food) error
	GetCache(ctx context.Context, ids []uint) ([]entity.CacheFood, error)
}
