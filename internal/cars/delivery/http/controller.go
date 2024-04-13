package http

import (
	"context"
	"fmt"
	"helloWRLDs/bookings/pkg/types"
	"helloWRLDs/bookings/pkg/web"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
)

func (d *CarsDeliveryImpl) AddCarController(w http.ResponseWriter, r *http.Request) {
	web.EncodeJson(w, 201, types.Response{Message: "Car added successfully"})
}

func (d *CarsDeliveryImpl) GetCarsController(w http.ResponseWriter, r *http.Request) {
	limit, _ := strconv.Atoi(r.URL.Query().Get("limit"))
	offset, _ := strconv.Atoi(r.URL.Query().Get("offset"))
	if limit <= 0 {
		limit = 10
	}
	if offset < 0 {
		offset = 0
	}
	ctx := context.WithValue(r.Context(), "limit", limit)
	ctx = context.WithValue(ctx, "offset", offset)
	cars, err := d.uc.GetCars(ctx)
	if err != nil {
		web.EncodeJson(w, 500, types.Response{Message: err.Error()})
		return
	}
	web.EncodeJson(w, 200, cars)
}

func (d *CarsDeliveryImpl) GetCarController(w http.ResponseWriter, r *http.Request) {
	web.EncodeJson(w, 200, types.Response{Message: fmt.Sprintf("Get Car with id=%s", chi.URLParam(r, "id"))})
}

func (d *CarsDeliveryImpl) DeleteCarController(w http.ResponseWriter, r *http.Request) {
	web.EncodeJson(w, 200, types.Response{Message: fmt.Sprintf("Delete Car with id=%s", chi.URLParam(r, "id"))})
}
