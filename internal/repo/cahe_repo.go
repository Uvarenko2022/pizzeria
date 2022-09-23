package repo

import (
	"Uvarenko2022/restaurant/database"
	"Uvarenko2022/restaurant/internal/entity"
	"context"
	"encoding/json"
	"log"
	"strconv"
)

type CacheRepo struct {
	*database.Redis
}

func (r *CacheRepo) UpdateCache(ctx context.Context, food []entity.Food) error {
	for _, v := range food {
		str, _ := json.Marshal(entity.FoodToCacheFood(v))
		result := r.Client.Set(ctx, strconv.Itoa(int(v.ID)), str, 0)

		if result.Err() != nil {
			log.Println(result.Err())
			return result.Err()
		}
	}

	return nil
}

func (r *CacheRepo) GetAllCache(ctx context.Context) ([]entity.CacheFood, error) {
	var food []entity.CacheFood
	var ids []int

	if err := r.Client.Keys(ctx, "*").ScanSlice(&ids); err != nil {
		log.Println(err)
		return nil, err
	}
	for i, v := range ids {
		result, _ := r.Client.Get(ctx, strconv.Itoa(v)).Result()
		food = append(food, entity.CacheFood{})

		if err := json.Unmarshal([]byte(result), &food[i]); err != nil {
			log.Println(err)
			return nil, err
		}
	}

	return food, nil
}

func (r *CacheRepo) GetProperCahce(ctx context.Context, ids []uint) ([]entity.CacheFood, error) {
	food := make([]entity.CacheFood, len(ids), len(ids))

	for i, v := range ids {
		str, _ := r.Client.Get(ctx, strconv.Itoa(int(v))).Result()
		err := json.Unmarshal([]byte(str), &food[i])

		if err != nil {
			log.Println(err)
			return nil, err
		}
	}

	return food, nil
}

func newCacheR(cache *database.Redis) ICacheRepo {
	return &CacheRepo{cache}
}
