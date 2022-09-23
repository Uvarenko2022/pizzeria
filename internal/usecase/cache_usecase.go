package usecase

import (
	"Uvarenko2022/restaurant/internal/entity"
	"Uvarenko2022/restaurant/internal/repo"
	"context"
)

type CacheUC struct {
	repo repo.ICacheRepo
}

func NewCacheUC(repo repo.ICacheRepo) *CacheUC {
	return &CacheUC{repo: repo}
}

func (uc *CacheUC) UpdateCache(ctx context.Context, food []entity.Food) error {
	return uc.repo.UpdateCache(ctx, food)
}

//Returns slice of food. If you an empty slice is passed returns all food
func (uc *CacheUC) GetCache(ctx context.Context, ids []uint) ([]entity.CacheFood, error) {
	if len(ids) == 0 {
		return uc.repo.GetAllCache(ctx)
	}

	return uc.repo.GetProperCahce(ctx, ids)
}
