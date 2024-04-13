package usecase

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	repo "helloWRLDs/bookings/internal/cars/repository"
	"helloWRLDs/bookings/pkg/types"
)

type CarsUseCase interface {
	GetCars(ctx context.Context) (types.PaginationResponse, error)
}

type CarsUseCaseImpl struct {
	repo repo.CarsRepository
}

func NewUseCase(db *sql.DB) *CarsUseCaseImpl {
	return &CarsUseCaseImpl{
		repo: repo.NewCarsRepository(db),
	}
}

func (u *CarsUseCaseImpl) GetCars(ctx context.Context) (types.PaginationResponse, error) {
	// Get vals from ctx
	limit, ok := ctx.Value("limit").(int)
	if !ok {
		return types.PaginationResponse{}, errors.New("query error")
	}
	offset, ok := ctx.Value("offset").(int)
	if !ok {
		return types.PaginationResponse{}, errors.New("query error")
	}
	url, ok := ctx.Value("url").(string)
	if !ok {
		return types.PaginationResponse{}, errors.New("query error")
	}

	// Get data from database
	cars, err := u.repo.GetAll(limit, offset)
	if err != nil {
		return types.PaginationResponse{}, err
	}

	// Set the response
	var response types.PaginationResponse
	response.Content = cars
	if offset < u.repo.Length() {
		response.Next = fmt.Sprintf("%s?limit=%d&offset=%d", url, limit, limit+offset)
	}
	if offset >= limit {
		response.Prev = fmt.Sprintf("%s?limit=%d&offset=%d", url, limit, offset-limit)
	}
	return response, nil
}
