package repository

import (
	"database/sql"
	"fmt"
	dm "helloWRLDs/bookings/internal/cars/domain"
)

type CarsRepositoryImpl struct {
	DB *sql.DB
}

func NewCarsRepository(db *sql.DB) *CarsRepositoryImpl {
	return &CarsRepositoryImpl{
		DB: db,
	}
}

func (r *CarsRepositoryImpl) Update(id int, c *dm.Car) (*dm.Car, error) {
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
	row := r.DB.QueryRow(
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
	if err := r.DB.QueryRow(stmt, id).Scan(&exist); err != nil {
		return false
	}
	return exist
}

func (r *CarsRepositoryImpl) Length() int {
	var length int
	stmt := `SELECT COUNT(*) FROM cars`
	if err := r.DB.QueryRow(stmt).Scan(&length); err != nil {
		return 0
	}
	return length
}

func (r *CarsRepositoryImpl) Get(id int) (dm.Car, error) {
	var c dm.Car
	stmt := `SELECT * FROM cars WHERE id=$1`
	err := r.DB.QueryRow(stmt, id).Scan(
		&c.ID, &c.Vendor, &c.Model, &c.Year, &c.IsAvailable, &c.Body, &c.EngineCapacity, &c.Mileage, &c.Color,
	)
	if err != nil {
		return dm.Car{}, err
	}
	return c, nil
}

func (r *CarsRepositoryImpl) Insert(c *dm.Car) (int, error) {
	stmt := `
	INSERT INTO cars(vendor, model, year, is_available, body, engine_capacity, mileage, color) 
	VALUES($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id;`
	var insertedId int
	row := r.DB.QueryRow(
		stmt, c.Vendor, c.Model, c.Year, c.IsAvailable, c.Body, c.EngineCapacity, c.Mileage, c.Color,
	)
	if err := row.Scan(&insertedId); err != nil {
		return 0, err
	}
	return insertedId, nil
}

func (r *CarsRepositoryImpl) GetAll(filters *dm.Filters) ([]dm.Car, error) {
	var cars []dm.Car
	stmt := fmt.Sprintf(`SELECT * FROM cars WHERE %s LIKE %s ORDER BY %s %s LIMIT $1 OFFSET $2`, filters.FilterType, filters.Filter, filters.Sort, filters.Order)
	rows, err := r.DB.Query(stmt, filters.Limit, filters.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var c dm.Car
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
	_, err := r.DB.Exec(stmt, id)
	if err != nil {
		return err
	}
	return nil
}
