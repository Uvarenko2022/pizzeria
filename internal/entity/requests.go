package entity

type OrderRequest struct {
	FoodIds []uint `json:"foodids"`
	State   State  `json:"state"`
	Comment string `json:"comment"`
}
