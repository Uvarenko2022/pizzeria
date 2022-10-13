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
	IIngUC
}

func New(repo *repo.Repository) *PiizzaUseCase {
	inuc := NewIngUC(repo.IIngRepo)
	fuc := NewFoodUC(repo.Ifoodr, inuc.GetIng)
	ouc := NewOrderUC(repo.Iordrrepo, fuc.GetFood)
	cuc := NewCacheUC(repo.Ichrepo)

	return &PiizzaUseCase{
		fuc,
		ouc,
		cuc,
		inuc,
	}
}

type IIngUC interface {
	AddIng(ing *entity.Ingridient) error
	DeleteIng(id uint) error
	UpdateIng(ing *entity.Ingridient) error
	GetIng() ([]entity.Ingridient, error)
}

type IFoodUC interface {
	GetFood(ids []uint) ([]entity.Food, error)
	AddFood(food *entity.FoodRequest) error
	UpdateFood(food *entity.Food) error
}

type IOrderUC interface {
	CreateOrder(order *entity.OrderRequest) error
	GetOrders(limit int, offset int) ([]entity.Order, error)
	UpdateOrder(order *entity.OrderRequest) error
}

type ICacheUC interface {
	UpdateCache(ctx context.Context, food []entity.Food) error
	GetCache(ctx context.Context, ids []uint) ([]entity.CacheFood, error)
}
