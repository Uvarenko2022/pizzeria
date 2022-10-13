package usecase

import (
	"Uvarenko2022/restaurant/internal/entity"
	"Uvarenko2022/restaurant/internal/repo"
)

type OrderUC struct {
	repo repo.IOrderRepo
	gf   func(ids []uint) ([]entity.Food, error)
}

func NewOrderUC(repo repo.IOrderRepo, gf func(ids []uint) ([]entity.Food, error)) *OrderUC {
	return &OrderUC{repo: repo, gf: gf}
}

func (app *OrderUC) CreateOrder(order *entity.OrderRequest) error {
	norder := &entity.Order{}
	food, err := app.gf(order.FoodIds)

	if err != nil {
		return err
	}

	for i, v := range food {
		norder.Food = append(norder.Food, &food[i])
		norder.TotalCost += v.Cost
	}

	norder.State = order.State
	norder.Comment = order.Comment

	return app.repo.CreateOrder(norder)
}

func (app *OrderUC) GetOrders(limit int, offset int) ([]entity.Order, error) {
	return app.repo.GetOrders(limit, offset)
}

func (app *OrderUC) UpdateOrder(order *entity.OrderRequest) error {
	norder := &entity.Order{}
	food, err := app.gf(order.FoodIds)

	if err != nil {
		return err
	}

	for i, v := range food {
		norder.Food = append(norder.Food, &food[i])
		norder.TotalCost += v.Cost
	}

	norder.State = order.State
	norder.Comment = order.Comment

	return app.repo.UpdateOrder(norder)
}
