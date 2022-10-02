package http

import (
	"Uvarenko2022/restaurant/internal/entity"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func (r *Rout) AddIngs(w http.ResponseWriter, req *http.Request) {
	var item *entity.Ingridient
	if err := json.NewDecoder(req.Body).Decode(item); err != nil {
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	if err := r.uc.IIngUC.AddIng(item); err != nil {
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (r *Rout) DeleteIng(w http.ResponseWriter, req *http.Request) {
	str := chi.URLParam(req, "id")
	id, err := strconv.Atoi(str)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	r.uc.IIngUC.DeleteIng(uint(id))
}

func (r *Rout) UpdateIng(w http.ResponseWriter, req *http.Request) {
	var item *entity.Ingridient
	if err := json.NewDecoder(req.Body).Decode(item); err != nil {
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	if err := r.uc.IIngUC.UpdateIng(item); err != nil {
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (r *Rout) GetIngs(w http.ResponseWriter, req *http.Request) {
	items, err := r.uc.GetIng()
	if err != nil {
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	json.NewEncoder(w).Encode(items)
}
