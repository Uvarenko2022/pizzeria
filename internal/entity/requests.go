package entity

type OrderRequest struct {
	FoodIds []uint `json:"foodids"`
	State   State  `json:"state"`
	Comment string `json:"comment"`
}

type FoodRequest struct {
	IngridientIds []uint    `json:"ingids"`
	Name          string    `json:"name"`
	Comment       string    `json:"comment"`
	Type          FoodType  `json:"foodtype"`
	Cost          float32   `json:"cost"`
	Mass          []float32 `json:"mass"`
}

type IngRequest struct {
}
