package repository

import (
	"database/sql"
	"helloWRLDs/bookings/internal/cars/domain"

	logger "github.com/sirupsen/logrus"
)

type CarsRepository interface {
	Get(id int64) (domain.Car, error)
	Insert(car domain.Car) (int64, error)
	GetAll(limit int, offset int) ([]domain.Car, error)
	Delete(id int64) error
}

type CarsRepositoryImpl struct {
	Db *sql.DB
}

func NewCarsRepository(db *sql.DB) *CarsRepositoryImpl {
	return &CarsRepositoryImpl{
		Db: db,
	}
}

func (r *CarsRepositoryImpl) Get(id int64) (domain.Car, error) {
	var c domain.Car
	stmt := `SELECT * FROM cars WHERE id=$1`
	err := r.Db.QueryRow(stmt, id).Scan(&c.ID, &c.Vendor, &c.Model, &c.Year, &c.IsAvailable, &c.Body)
	if err != nil {
		return domain.Car{}, err
	}
	return c, nil
}

func (r *CarsRepositoryImpl) Insert(car domain.Car) (int64, error) {
	return 1, nil
}

func (r *CarsRepositoryImpl) GetAll(limit int, offset int) ([]domain.Car, error) {
	var cars []domain.Car
	stmt := `SELECT * FROM cars ORDER BY id LIMIT $1 OFFSET $2`
	rows, err := r.Db.Query(stmt, limit, offset)
	if err != nil {
		logger.Error(err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var car domain.Car
		err := rows.Scan(&car.ID, &car.Vendor, &car.Model, &car.Year, &car.IsAvailable, &car.Body)
		if err != nil {
			return nil, err
		}
		cars = append(cars, car)
	}
	return cars, nil
}

func (r *CarsRepositoryImpl) Delete(id int64) error {
	return nil
}
