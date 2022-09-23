package entity

import (
	"gorm.io/gorm"
)

type FoodType int

const (
	drink FoodType = iota
	food
)

type Size int

const (
	Small Size = iota
	Middle
	Large
)

type State int

const (
	WaitingForCook State = iota
	Cooking
	Cooked
)

type Food struct {
	gorm.Model
	Type FoodType `json:"type" validate:"required,type"`
	Name string   `json:"name" validate:"required,name"`
	Cost float32  `json:"cost" validate:"required,cost"`
}

func FoodToCacheFood(food Food) *CacheFood {
	return &CacheFood{
		ID:   food.ID,
		Name: food.Name,
		Cost: food.Cost,
	}
}

type CacheFood struct {
	ID   uint
	Name string  `json:"name"`
	Cost float32 `json:"cost"`
}

type Order struct {
	gorm.Model
	FoodIds   []uint  `json:"foodids" gorm:"-:all;" validate:"required,foodids"`
	Food      []*Food `gorm:"many2many:food_order;"`
	State     State   `json:"state" gorm:"embedded" validate:"required,state"`
	TotalCost float32
}
