package http

import (
	"database/sql"
	"helloWRLDs/bookings/internal/cars/usecase"
)

type CarsDeliveryImpl struct {
	uc usecase.CarsUseCase
}

func NewCarsDelivery(db *sql.DB) *CarsDeliveryImpl {
	return &CarsDeliveryImpl{
		uc: usecase.NewUseCase(db),
	}
}
