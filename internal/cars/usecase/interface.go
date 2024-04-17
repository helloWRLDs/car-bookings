package usecase

import (
	"context"
	"helloWRLDs/bookings/internal/cars/domain"
	"helloWRLDs/bookings/pkg/types"
	ctx "helloWRLDs/bookings/pkg/types/context"
	rsp "helloWRLDs/bookings/pkg/types/responses"
)

type CarsUseCase interface {
	GetCars(ctx ctx.Context, filters domain.Filters) (rsp.Pagination, error)
	GetCar(ctx context.Context, id int) (*domain.Car, *types.Error)
	InsertCar(ctx context.Context, car *domain.Car) (int, *types.Error)
	DeleteCar(ctx context.Context, id int) *types.Error
	UpdateCar(ctx context.Context, id int, car *domain.Car) (*domain.Car, *types.Error)
}
