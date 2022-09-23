package http

import (
	"encoding/json"
	"log"
	"net/http"
)

//Cache
func (r *Rout) GetCache(w http.ResponseWriter, req *http.Request) {
	result, err := r.uc.GetCache(req.Context(), []uint{})

	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(result)
}
