package http

import (
	"Uvarenko2022/restaurant/internal/entity"
	"encoding/json"
	"log"
	"net/http"
)

func (r *Rout) AddFood(w http.ResponseWriter, req *http.Request) {
	item := &entity.Food{}
	json.NewDecoder(req.Body).Decode(item)

	if err := r.trans.Struct(item); err != nil {
		errs := r.trans.TranslateError(err)
		log.Println(errs)
		http.Error(w, errs, http.StatusBadRequest)
		return
	}

	if err := r.uc.IFoodUC.AddFood(item); err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (r *Rout) GetFood(w http.ResponseWriter, req *http.Request) {
	result, err := r.uc.IFoodUC.GetFood([]uint{})
	if err != nil {
		errs := r.trans.TranslateError(err)
		log.Println(errs)
		http.Error(w, errs, http.StatusInternalServerError)
	}

	json.NewEncoder(w).Encode(result)
}

func (r *Rout) UpdateFood(w http.ResponseWriter, req *http.Request) {
	item := &entity.Food{}
	json.NewDecoder(req.Body).Decode(item)

	if err := r.trans.Struct(item); err != nil {
		errs := r.trans.TranslateError(err)
		log.Println(errs)
		http.Error(w, errs, http.StatusBadRequest)
		return
	}

	if err := r.uc.IFoodUC.UpdateFood(item); err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
