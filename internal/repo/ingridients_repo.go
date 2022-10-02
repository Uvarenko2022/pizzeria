package repo

import (
	"Uvarenko2022/restaurant/database"
	"Uvarenko2022/restaurant/internal/entity"
)

type IngRepo struct {
	*database.Postgre
	Model *entity.Ingridient
}

func newIngR(db *database.Postgre) IIngRepo {
	return &IngRepo{
		db,
		&entity.Ingridient{},
	}
}

func (r *IngRepo) AddIng(ing *entity.Ingridient) error {
	err := r.DB.Model(r.Model).Create(ing).Error

	return err
}

func (r *IngRepo) DeleteIng(id uint) error {
	err := r.DB.Where("id = ?", id).Delete(r.Model).Error

	return err
}

func (r *IngRepo) UpdateIng(ing *entity.Ingridient) error {
	err := r.DB.Model(r.Model).Where("id = ?", ing.ID).Updates(ing).Error

	return err
}

func (r *IngRepo) GetIng() ([]entity.Ingridient, error) {
	var ings []entity.Ingridient
	err := r.DB.Model(r.Model).Find(&ings).Error

	return ings, err
}
