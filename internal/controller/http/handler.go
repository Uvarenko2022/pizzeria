package http

import (
	"Uvarenko2022/restaurant/internal/usecase"
	"Uvarenko2022/restaurant/internal/validate"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Rout struct {
	uc    *usecase.PiizzaUseCase
	trans validate.Translate
}

func New(uc *usecase.PiizzaUseCase, tr validate.Translate) *Rout {
	return &Rout{uc, tr}
}

func RegisterRouts(rout *Rout) {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	//food
	r.Route("/food", func(router chi.Router) {
		router.Post("/add", rout.AddFood)
		router.Get("/get", rout.GetFood)
		router.Patch("/update", rout.UpdateFood)
	})

	//cache
	r.Get("/cache/get", rout.GetCache)

	//order
	r.Route("/order", func(router chi.Router) {
		router.Post("/create", rout.CreateOrder)
		router.Patch("/update", rout.UpdateOrder)
		router.Get("/limit={limit}offset={offset}", rout.GetOrders)
	})

	//ingridients
	r.Route("/ingridient", func(router chi.Router) {
		router.Get("/get", rout.GetIngs)
		router.Post("/add", rout.AddIngs)
		router.Patch("/update", rout.UpdateIng)
		router.Delete("/delete-{id}", rout.DeleteIng)
	})

	http.ListenAndServe(":8080", r)
}
