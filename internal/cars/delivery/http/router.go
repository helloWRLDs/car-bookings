package http

import (
	"github.com/go-chi/chi"
)

func (d *CarsDeliveryImpl) Routes() chi.Router {
	r := chi.NewRouter()

	r.Get("/", d.GetCarsController)
	r.Get("/{id}", d.GetCarController)
	r.Post("/", d.AddCarController)
	r.Delete("/{id}", d.DeleteCarController)

	return r
}
