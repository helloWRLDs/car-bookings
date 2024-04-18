package http

import (
	"fmt"
	dm "helloWRLDs/bookings/internal/cars/domain"
	ctx "helloWRLDs/bookings/pkg/types/context"
	resp "helloWRLDs/bookings/pkg/types/responses"
	enc "helloWRLDs/bookings/pkg/web/encoders"
	errs "helloWRLDs/bookings/pkg/web/errors"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
)

func (d *CarsDeliveryImpl) AddCarController(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	car, err := enc.DecodeJson[dm.Car](r)
	if err != nil {
		enc.SendErr(w, errs.ErrUnpocessableEntity)
		return
	}
	id, error := d.uc.InsertCar(ctx, &car)
	if error != nil {
		enc.SendErr(w, error)
		return
	}
	enc.EncodeJson(w, 201, resp.Message{Message: fmt.Sprintf("inserted car with id=%d", id)})
}

func (d *CarsDeliveryImpl) GetCarsController(w http.ResponseWriter, r *http.Request) {
	ctx := ctx.Context{
		Context: r.Context(),
		Data:    make(map[string]interface{}, 1),
	}
	filters := dm.ExtractFilters(r)
	ctx.Data["url"] = fmt.Sprintf("http://%s%s", r.Host, r.URL.Path)
	response, error := d.uc.GetCars(ctx, filters)
	if error != nil {
		enc.SendErr(w, error)
		return
	}
	enc.EncodeJson(w, 200, response)
}

func (d *CarsDeliveryImpl) GetCarController(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		enc.SendErr(w, errs.ErrBadRequest)
		return
	}
	ctx := r.Context()
	car, error := d.uc.GetCar(ctx, id)
	if error != nil {
		enc.SendErr(w, error)
		return
	}
	enc.EncodeJson(w, 200, car)
}

func (d *CarsDeliveryImpl) DeleteCarController(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		enc.SendErr(w, errs.ErrBadRequest)
		return
	}
	if error := d.uc.DeleteCar(ctx, id); error != nil {
		enc.SendErr(w, error)
		return
	}
	enc.EncodeJson(w, 204, nil)
}

func (d *CarsDeliveryImpl) UpdateCarController(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		enc.SendErr(w, errs.ErrBadRequest)
		return
	}
	car, err := enc.DecodeJson[dm.Car](r)
	if err != nil {
		enc.SendErr(w, errs.ErrUnpocessableEntity)
		return
	}
	newCar, error := d.uc.UpdateCar(ctx, id, &car)
	if error != nil {
		enc.SendErr(w, error)
		return
	}
	enc.EncodeJson(w, 200, newCar)
}
