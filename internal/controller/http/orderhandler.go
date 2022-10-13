package http

import (
	"Uvarenko2022/restaurant/internal/entity"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
)

// Order
func (r *Rout) CreateOrder(w http.ResponseWriter, req *http.Request) {
	item := &entity.OrderRequest{}
	if err := json.NewDecoder(req.Body).Decode(item); err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	//validating
	if err := r.trans.Struct(item); err != nil {
		errs := r.trans.TranslateError(err)
		log.Println(errs)
		http.Error(w, errs, http.StatusBadRequest)
		return
	}

	if err := r.uc.CreateOrder(item); err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (r *Rout) UpdateOrder(w http.ResponseWriter, req *http.Request) {
	item := &entity.OrderRequest{}
	if err := json.NewDecoder(req.Body).Decode(item); err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := r.uc.UpdateOrder(item); err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (r *Rout) GetOrders(w http.ResponseWriter, req *http.Request) {
	slimit := chi.URLParam(req, "limit")
	soffset := chi.URLParam(req, "offset")

	limit, _ := strconv.Atoi(slimit)
	offset, _ := strconv.Atoi(soffset)

	result, err := r.uc.GetOrders(limit, offset)

	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(result); err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}
