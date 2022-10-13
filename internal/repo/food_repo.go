package repo

import (
	"Uvarenko2022/restaurant/database"
	"Uvarenko2022/restaurant/internal/entity"
	"gorm.io/gorm"
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

func (r *FoodRepo) Put(f *entity.Food, masses []float32) error {
	err := r.DB.Transaction(func(tx *gorm.DB) error {
		err := r.DB.Create(f).Error
		if err != nil {
			return err
		}

		for i, v := range f.Ingridients {
			err := r.DB.Model(&entity.FoodIng{}).
				Where("food_id = ? AND ingridient_id = ?", f.ID, v.ID).
				Update("mass", masses[i]).Error
			if err != nil {
				return err
			}
		}
		return nil
	})

	return err
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
