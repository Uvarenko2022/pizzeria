package repo

import (
	"Uvarenko2022/restaurant/database"
	"Uvarenko2022/restaurant/internal/entity"
	"log"
)

type OrderRepo struct {
	*database.Postgre
	Model *entity.Order
}

func (r *OrderRepo) CreateOrder(order *entity.Order) error {
	if err := r.DB.Model(r.Model).Create(order).Error; err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (r *OrderRepo) GetOrders(limit int, offset int) ([]entity.Order, error) {
	result := make([]entity.Order, limit)

	if err := r.DB.Limit(limit).Offset(offset).Find(&result).Error; err != nil {
		log.Println(err)
		return nil, err
	}

	return result, nil
}

func (r *OrderRepo) UpdateOrder(order *entity.Order) error {
	if err := r.DB.Model(r.Model).Where("ID = ?", order.ID).Updates(order).Error; err != nil {
		log.Println(err)
		return err
	}

	if err := r.DB.Model(r.Model).Association("Food").Replace(order.Food); err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func newOrderR(db *database.Postgre) IOrderRepo {
	return &OrderRepo{
		db,
		&entity.Order{},
	}
}
