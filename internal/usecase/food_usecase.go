package usecase

import (
	"Uvarenko2022/restaurant/internal/entity"
	"Uvarenko2022/restaurant/internal/repo"
)

type FoodUC struct {
	repo repo.IFoodRepo
}

func NewFoodUC(repo repo.IFoodRepo) *FoodUC{
	return &FoodUC{repo: repo}
}

//Returns slice of food with passed ids. If an empty slice is passed than returns all food.
func (app *FoodUC) GetFood(ids []uint) ([]entity.Food, error) {
	if len(ids) == 0 {
		return app.repo.GetAllFood()
	}

	return app.repo.GetProperFood(ids)
}

func (app *FoodUC) AddFood(food *entity.Food) error{
	return app.repo.Put(food)
}

func (app *FoodUC) UpdateFood(food *entity.Food) error{
	return app.repo.UpdateFood(food)
}
