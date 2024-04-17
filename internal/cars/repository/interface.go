package repository

import (
	"helloWRLDs/bookings/internal/cars/domain"
)

type CarsRepository interface {
	Get(id int) (domain.Car, error)
	Insert(car *domain.Car) (int, error)
	GetAll(filters domain.Filters) ([]domain.Car, error)
	Delete(id int) error
	Length() int
	Exists(id int) bool
	Update(id int, c *domain.Car) (*domain.Car, error)
}
