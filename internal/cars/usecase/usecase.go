package usecase

import (
	"context"
	"database/sql"
	"fmt"
	"helloWRLDs/bookings/internal/cars/domain"
	repo "helloWRLDs/bookings/internal/cars/repository"
	"helloWRLDs/bookings/pkg/types"
	ctx "helloWRLDs/bookings/pkg/types/context"
	rsp "helloWRLDs/bookings/pkg/types/responses"
)

type CarsUseCaseImpl struct {
	repo repo.CarsRepository
}

func NewUseCase(db *sql.DB) *CarsUseCaseImpl {
	return &CarsUseCaseImpl{
		repo: repo.NewCarsRepository(db),
	}
}

func (u *CarsUseCaseImpl) UpdateCar(ctx context.Context, id int, car *domain.Car) (*domain.Car, *types.Error) {
	if !u.repo.Exists(id) {
		return nil, types.NewErr(fmt.Sprintf("object with id=%d does not exist", id), 404)
	}
	updatedCar, err := u.repo.Update(id, car)
	if err != nil {
		return nil, types.NewErr(err.Error(), 500)
	}
	return updatedCar, nil
}

func (u *CarsUseCaseImpl) InsertCar(ctx context.Context, car *domain.Car) (int, *types.Error) {
	if err := car.Validate(); err != nil {
		return 0, types.NewErr(err.Error(), 403)
	}
	id, error := u.repo.Insert(car)
	if error != nil {
		return 0, types.NewErr(error.Error(), 500)
	}
	return id, nil
}

func (u *CarsUseCaseImpl) GetCar(ctx context.Context, id int) (*domain.Car, *types.Error) {
	if !u.repo.Exists(id) {
		return nil, types.NewErr(fmt.Sprintf("object with id=%d not found", id), 404)
	}
	car, err := u.repo.Get(id)
	if err != nil {
		return nil, types.NewErr(err.Error(), 500)
	}
	return &car, nil
}

func (u *CarsUseCaseImpl) GetCars(ctx ctx.Context, filters domain.Filters) (rsp.Pagination, error) {
	// Get data from database
	cars, err := u.repo.GetAll(filters)
	if err != nil {
		return rsp.Pagination{}, err
	}

	// Set the response
	var response rsp.Pagination
	response.Content = cars
	if filters.Offset < u.repo.Length() {
		response.Next = fmt.Sprintf("%s?limit=%d&offset=%d", ctx.Data["url"], filters.Limit, filters.Limit+filters.Offset)
	}
	if filters.Offset >= filters.Limit {
		response.Prev = fmt.Sprintf("%s?limit=%d&offset=%d", ctx.Data["url"], filters.Limit, filters.Offset-filters.Limit)
	} else if filters.Offset > 0 {
		response.Prev = fmt.Sprintf("%s?limit=%d&offset=%d", ctx.Data["url"], filters.Limit, 0)
	}
	return response, nil
}

func (u *CarsUseCaseImpl) DeleteCar(ctx context.Context, id int) *types.Error {
	if id < 0 {
		return types.NewErr("wrong id format", 400)
	}
	if !u.repo.Exists(id) {
		return types.NewErr(fmt.Sprintf("object with id=%d not found", id), 404)
	}
	if err := u.repo.Delete(id); err != nil {
		return types.NewErr(err.Error(), 500)
	}
	return nil
}
