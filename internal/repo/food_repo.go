package repo

import (
	"Uvarenko2022/restaurant/database"
	"Uvarenko2022/restaurant/internal/entity"
)

type FoodRepo struct {
	*database.Postgre
	Model *entity.Food
}

func (r *FoodRepo) GetProperFood(ids []uint) ([]entity.Food, error) {
	items := make([]entity.Food, len(ids), len(ids))
	result := r.DB.Where("id in ?", ids).Find(&items)

	return items, result.Error
}

func (r *FoodRepo) GetAllFood() ([]entity.Food, error) {
	var items []entity.Food
	result := r.DB.Find(&items)
	return items, result.Error
}

func (r *FoodRepo) Put(f *entity.Food) error {
	result := r.DB.Create(f)

	return result.Error
}

func (r *FoodRepo) UpdateFood(food *entity.Food) error {
	result := r.DB.Where("id = ?", food.ID).Find(food)
	return result.Error
}

func newFoodR(db *database.Postgre) IFoodRepo {
	return &FoodRepo{
		db,
		&entity.Food{},
	}
}
