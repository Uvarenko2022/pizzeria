package usecase

import (
	"Uvarenko2022/restaurant/internal/entity"
	"Uvarenko2022/restaurant/internal/repo"
)

type OrderUC struct {
	repo repo.IOrderRepo
}

func NewOrderUC(repo repo.IOrderRepo) *OrderUC {
	return &OrderUC{repo: repo}
}

func (app *OrderUC) CreateOrder(order *entity.Order, food []entity.Food) error{
	for i := range food {
		order.Food = append(order.Food, &food[i])
	}

	return app.repo.CreateOrder(order)
}

func (app *OrderUC) GetOrders(limit int, offset int) ([]entity.Order, error) {
	return app.repo.GetOrders(limit, offset)
}

func (app *OrderUC) UpdateOrder(order *entity.Order, food []entity.Food) error{
	for _, v := range food {
		order.Food = append(order.Food, &v)
	}

	return app.repo.UpdateOrder(order)
}
