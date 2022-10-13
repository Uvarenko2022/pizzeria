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

type FoodIng struct {
	FoodID       uint
	IngridientId uint
	Mass         float32
}

type Food struct {
	gorm.Model
	Type        FoodType      `json:"type" validate:"required,type"`
	Name        string        `json:"name" validate:"required,name"`
	Cost        float32       `json:"cost" validate:"required,cost"`
	Comment     string        `json:"comment"`
	Ingridients []*Ingridient `gorm:"many2many:food_ingridients;"`
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
	Food      []*Food `gorm:"many2many:food_order;"`
	State     State   `json:"state" gorm:"embedded" validate:"required,state"`
	TotalCost float32
	Comment   string `json:"comment"`
}

type Ingridient struct {
	gorm.Model
	CurrentAmount float32 `json:"currentamount"`
	Name          string  `json:"name"`
	Comment       string  `json:"comment"`
}
