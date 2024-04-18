package usecase

import (
	"context"
	"database/sql"
	"fmt"
	dm "helloWRLDs/bookings/internal/cars/domain"
	repo "helloWRLDs/bookings/internal/cars/repository"
	ctx "helloWRLDs/bookings/pkg/types/context"
	resp "helloWRLDs/bookings/pkg/types/responses"
	errs "helloWRLDs/bookings/pkg/web/errors"

	logger "github.com/sirupsen/logrus"
)

type CarsUseCaseImpl struct {
	repo repo.CarsRepository
}

func NewUseCase(db *sql.DB) *CarsUseCaseImpl {
	return &CarsUseCaseImpl{
		repo: repo.NewCarsRepository(db),
	}
}

func (u *CarsUseCaseImpl) UpdateCar(ctx context.Context, id int, car *dm.Car) (*dm.Car, *errs.Error) {
	if !u.repo.Exists(id) {
		return nil, errs.ErrNotFound
	}
	updatedCar, err := u.repo.Update(id, car)
	if err != nil {
		logger.WithField("err", err.Error()).Error("SQL Query Error")
		return nil, errs.ErrInternal
	}
	return updatedCar, nil
}

func (u *CarsUseCaseImpl) InsertCar(ctx context.Context, car *dm.Car) (int, *errs.Error) {
	if err := car.Validate(); err != nil {
		return 0, errs.ErrUnpocessableEntity
	}
	id, error := u.repo.Insert(car)
	if error != nil {
		logger.WithField("err", error.Error()).Error("SQL Query Error")
		return 0, errs.ErrInternal
	}
	return id, nil
}

func (u *CarsUseCaseImpl) GetCar(ctx context.Context, id int) (*dm.Car, *errs.Error) {
	if !u.repo.Exists(id) {
		return nil, errs.ErrNotFound
	}
	car, err := u.repo.Get(id)
	if err != nil {
		logger.WithField("err", err.Error()).Error("SQL Query Error")
		return nil, errs.ErrInternal
	}
	return &car, nil
}

func (u *CarsUseCaseImpl) GetCars(ctx ctx.Context, filters dm.Filters) (*resp.Pagination, *errs.Error) {
	cars, err := u.repo.GetAll(filters)
	if err != nil {
		logger.WithField("err", err.Error()).Error("SQL Query Error")
		return nil, errs.ErrInternal
	}

	response := &resp.Pagination{}
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

func (u *CarsUseCaseImpl) DeleteCar(ctx context.Context, id int) *errs.Error {
	if id < 0 {
		return errs.ErrBadRequest
	}
	if !u.repo.Exists(id) {
		return errs.ErrNotFound
	}
	if err := u.repo.Delete(id); err != nil {
		logger.WithField("err", err.Error()).Error("SQL Query Error")
		return errs.ErrInternal
	}
	return nil
}
