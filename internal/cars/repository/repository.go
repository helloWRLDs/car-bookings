package repository

import (
	"database/sql"
	"fmt"
	"helloWRLDs/bookings/internal/cars/domain"

	logger "github.com/sirupsen/logrus"
)

type CarsRepositoryImpl struct {
	Db *sql.DB
}

func NewCarsRepository(db *sql.DB) *CarsRepositoryImpl {
	return &CarsRepositoryImpl{
		Db: db,
	}
}

func (r *CarsRepositoryImpl) Update(id int, c *domain.Car) (*domain.Car, error) {
	stmt := `UPDATE cars
			SET vendor=$1, 
				model=$2, 
				year=$3, 
				is_available=$4, 
				body=$5, 
				engine_capacity=$6, 
				mileage=$7, 
				color=$8
			WHERE id=$9 
			RETURNING *;`
	row := r.Db.QueryRow(
		stmt, c.Vendor, c.Model, c.Year, c.IsAvailable, c.Body, c.EngineCapacity, c.Mileage, c.Color, id,
	)
	err := row.Scan(
		&c.ID, &c.Vendor, &c.Model, &c.Year, &c.IsAvailable, &c.Body, &c.EngineCapacity, &c.Mileage, &c.Color,
	)
	if err != nil {
		return nil, err
	}
	return c, nil
}

func (r *CarsRepositoryImpl) Exists(id int) bool {
	var exist bool
	stmt := `SELECT EXISTS(SELECT TRUE FROM cars WHERE id=$1)`
	if err := r.Db.QueryRow(stmt, id).Scan(&exist); err != nil {
		return false
	}
	return exist
}

func (r *CarsRepositoryImpl) Length() int {
	var length int
	stmt := `SELECT COUNT(*) FROM cars`
	if err := r.Db.QueryRow(stmt).Scan(&length); err != nil {
		return 0
	}
	return length
}

func (r *CarsRepositoryImpl) Get(id int) (domain.Car, error) {
	var c domain.Car
	stmt := `SELECT * FROM cars WHERE id=$1`
	err := r.Db.QueryRow(stmt, id).Scan(
		&c.ID, &c.Vendor, &c.Model, &c.Year, &c.IsAvailable, &c.Body, &c.EngineCapacity, &c.Mileage, &c.Color,
	)
	if err != nil {
		return domain.Car{}, err
	}
	return c, nil
}

func (r *CarsRepositoryImpl) Insert(c *domain.Car) (int, error) {
	stmt := `
	INSERT INTO cars(vendor, model, year, is_available, body, engine_capacity, mileage, color) 
	VALUES($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id;`
	var insertedId int
	row := r.Db.QueryRow(
		stmt, c.Vendor, c.Model, c.Year, c.IsAvailable, c.Body, c.EngineCapacity, c.Mileage, c.Color,
	)
	if err := row.Scan(&insertedId); err != nil {
		return 0, err
	}
	return insertedId, nil
}

func (r *CarsRepositoryImpl) GetAll(filters domain.Filters) ([]domain.Car, error) {
	fmt.Println(filters.Order)
	var cars []domain.Car
	stmt := fmt.Sprintf(`SELECT * FROM cars WHERE %s LIKE '%%%s%%' ORDER BY $1 %s LIMIT $2 OFFSET $3`, filters.FilterType, filters.Filter, filters.Order)
	// if filters.Order == "DESC" {
	// 	stmt += "DESC LIMIT $2 OFFSET $3"
	// } else {
	// 	stmt += "ASC LIMIT $2 OFFSET $3"
	// }
	fmt.Println(stmt)
	rows, err := r.Db.Query(stmt, filters.Sort, filters.Limit, filters.Offset)
	if err != nil {
		logger.Error(err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var c domain.Car
		err := rows.Scan(
			&c.ID, &c.Vendor, &c.Model, &c.Year, &c.IsAvailable, &c.Body, &c.EngineCapacity, &c.Mileage, &c.Color,
		)
		if err != nil {
			return nil, err
		}
		cars = append(cars, c)
	}
	return cars, nil
}

func (r *CarsRepositoryImpl) Delete(id int) error {
	stmt := `DELETE FROM cars WHERE id=$1`
	_, err := r.Db.Exec(stmt, id)
	if err != nil {
		return err
	}
	return nil
}
