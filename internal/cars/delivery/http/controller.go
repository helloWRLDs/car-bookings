package http

import (
	"fmt"
	"helloWRLDs/bookings/internal/cars/domain"
	ctx "helloWRLDs/bookings/pkg/types/context"
	rsp "helloWRLDs/bookings/pkg/types/responses"
	"helloWRLDs/bookings/pkg/web"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
)

func (d *CarsDeliveryImpl) AddCarController(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	car, err := web.DecodeJson[domain.Car](r)
	if err != nil {
		web.EncodeJson(w, 403, rsp.Message{Message: "unpocessable entity"})
		return
	}
	id, error := d.uc.InsertCar(ctx, &car)
	if error != nil {
		web.EncodeJson(w, error.Code(), rsp.Message{Message: error.Error()})
		return
	}
	web.EncodeJson(w, 201, rsp.Message{Message: fmt.Sprintf("inserted car with id=%d", id)})
}

func (d *CarsDeliveryImpl) GetCarsController(w http.ResponseWriter, r *http.Request) {
	ctx := ctx.Context{
		Context: r.Context(),
		Data:    make(map[string]interface{}, 1),
	}
	filters := domain.ExtractFilters(r)
	ctx.Data["url"] = fmt.Sprintf("http://%s%s", r.Host, r.URL.Path)

	response, err := d.uc.GetCars(ctx, filters)
	if err != nil {
		web.EncodeJson(w, 500, rsp.Message{})
		return
	}
	web.EncodeJson(w, 200, response)
}

func (d *CarsDeliveryImpl) GetCarController(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		web.EncodeJson(w, 400, rsp.Message{Message: "Bad Gateway"})
		return
	}
	ctx := r.Context()
	car, error := d.uc.GetCar(ctx, id)
	if error != nil {
		web.EncodeJson(w, error.Code(), rsp.Message{Message: error.Error()})
		return
	}
	web.EncodeJson(w, 200, car)
}

func (d *CarsDeliveryImpl) DeleteCarController(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		web.EncodeJson(w, 400, rsp.Message{Message: "bad gateway"})
		return
	}
	if error := d.uc.DeleteCar(ctx, id); error != nil {
		web.EncodeJson(w, error.Code(), error.Error())
		return
	}
	web.EncodeJson(w, 204, "")
}

func (d *CarsDeliveryImpl) UpdateCarController(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		web.EncodeJson(w, 400, rsp.Message{Message: "bad gateway"})
		return
	}
	car, err := web.DecodeJson[domain.Car](r)
	if err != nil {
		web.EncodeJson(w, 403, rsp.Message{Message: "unpocessable entity"})
		return
	}
	newCar, error := d.uc.UpdateCar(ctx, id, &car)
	if error != nil {
		web.EncodeJson(w, error.Code(), rsp.Message{Message: error.Error()})
		return
	}
	web.EncodeJson(w, 200, newCar)
}
