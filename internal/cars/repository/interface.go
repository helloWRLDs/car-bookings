package repository

import (
	dm "helloWRLDs/bookings/internal/cars/domain"
)

type CarsRepository interface {
	Get(id int) (dm.Car, error)
	Insert(car *dm.Car) (int, error)
	GetAll(filters dm.Filters) ([]dm.Car, error)
	Delete(id int) error
	Length() int
	Exists(id int) bool
	Update(id int, c *dm.Car) (*dm.Car, error)
}
