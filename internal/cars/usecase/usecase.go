package usecase

import (
	"context"
	"database/sql"
	"errors"
	"helloWRLDs/bookings/internal/cars/domain"
	repo "helloWRLDs/bookings/internal/cars/repository"
)

type CarsUseCase interface {
	GetCars(ctx context.Context) ([]domain.Car, error)
}

type CarsUseCaseImpl struct {
	repo repo.CarsRepository
}

func NewUseCase(db *sql.DB) *CarsUseCaseImpl {
	return &CarsUseCaseImpl{
		repo: repo.NewCarsRepository(db),
	}
}

func (u *CarsUseCaseImpl) GetCars(ctx context.Context) ([]domain.Car, error) {

	limit, ok := ctx.Value("limit").(int)
	if !ok {
		return nil, errors.New("query error")
	}
	offset, ok := ctx.Value("offset").(int)
	if !ok {
		return nil, errors.New("query error")
	}

	cars, err := u.repo.GetAll(limit, offset)
	if err != nil {
		return nil, err
	}
	return cars, nil
}
