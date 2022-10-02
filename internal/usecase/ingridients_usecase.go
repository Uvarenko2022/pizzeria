package usecase

import (
	"Uvarenko2022/restaurant/internal/entity"
	"Uvarenko2022/restaurant/internal/repo"
)

type IngUC struct {
	repo repo.IIngRepo
}

func NewIngUC(repo repo.IIngRepo) IIngUC {
	return &IngUC{
		repo: repo,
	}
}

func (uc *IngUC) AddIng(ing *entity.Ingridient) error {
	return uc.repo.AddIng(ing)
}

func (uc *IngUC) DeleteIng(id uint) error {
	return uc.repo.DeleteIng(id)
}

func (uc *IngUC) UpdateIng(ing *entity.Ingridient) error {
	return uc.repo.UpdateIng(ing)
}

func (uc *IngUC) GetIng() ([]entity.Ingridient, error) {
	return uc.repo.GetIng()
}
