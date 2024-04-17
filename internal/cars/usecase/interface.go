package usecase

import (
	"context"
	dm "helloWRLDs/bookings/internal/cars/domain"
	"helloWRLDs/bookings/pkg/types"
	ctx "helloWRLDs/bookings/pkg/types/context"
	rsp "helloWRLDs/bookings/pkg/types/responses"
)

type CarsUseCase interface {
	GetCars(ctx ctx.Context, filters dm.Filters) (rsp.Pagination, error)
	GetCar(ctx context.Context, id int) (*dm.Car, *types.Error)
	InsertCar(ctx context.Context, car *dm.Car) (int, *types.Error)
	DeleteCar(ctx context.Context, id int) *types.Error
	UpdateCar(ctx context.Context, id int, car *dm.Car) (*dm.Car, *types.Error)
}
