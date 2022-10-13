package usecase

import (
	"Uvarenko2022/restaurant/internal/entity"
	"Uvarenko2022/restaurant/internal/repo"
)

type FoodUC struct {
	repo repo.IFoodRepo
	gi   func() ([]entity.Ingridient, error)
}

func NewFoodUC(repo repo.IFoodRepo, gi func() ([]entity.Ingridient, error)) *FoodUC {
	return &FoodUC{repo: repo, gi: gi}
}

// Returns slice of food with passed ids. If an empty slice is passed than returns all food.
func (app *FoodUC) GetFood(ids []uint) ([]entity.Food, error) {
	if len(ids) == 0 {
		return app.repo.GetAllFood()
	}

	return app.repo.GetProperFood(ids)
}

func (app *FoodUC) AddFood(food *entity.FoodRequest) error {
	f := &entity.Food{}
	ings, err := app.gi()

	if err != nil {
		return err
	}

	for i, v := range food.IngridientIds {
		if v == ings[i].ID {
			f.Ingridients = append(f.Ingridients, &ings[i])
		}
	}

	f.Name = food.Name
	f.Cost = food.Cost
	f.Comment = food.Comment
	f.Type = food.Type

	return app.repo.Put(f, food.Mass)
}

func (app *FoodUC) UpdateFood(food *entity.Food) error {
	return app.repo.UpdateFood(food)
}
