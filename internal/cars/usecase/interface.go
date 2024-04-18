package usecase

import (
	"context"
	dm "helloWRLDs/bookings/internal/cars/domain"
	ctx "helloWRLDs/bookings/pkg/types/context"
	rsp "helloWRLDs/bookings/pkg/types/responses"
	errs "helloWRLDs/bookings/pkg/web/errors"
)

type CarsUseCase interface {
	GetCars(ctx ctx.Context, filters dm.Filters) (*rsp.Pagination, *errs.Error)
	GetCar(ctx context.Context, id int) (*dm.Car, *errs.Error)
	InsertCar(ctx context.Context, car *dm.Car) (int, *errs.Error)
	DeleteCar(ctx context.Context, id int) *errs.Error
	UpdateCar(ctx context.Context, id int, car *dm.Car) (*dm.Car, *errs.Error)
}
