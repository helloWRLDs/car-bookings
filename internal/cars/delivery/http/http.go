package http

import (
	"database/sql"
	uc "helloWRLDs/bookings/internal/cars/usecase"
)

type CarsDeliveryImpl struct {
	uc uc.CarsUseCase
}

func NewCarsDelivery(db *sql.DB) *CarsDeliveryImpl {
	return &CarsDeliveryImpl{
		uc: uc.NewUseCase(db),
	}
}
