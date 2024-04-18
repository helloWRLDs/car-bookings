package domain

import "fmt"

type Car struct {
	ID             int64   `json:"id,omitempty"`
	Vendor         string  `json:"vendor"`
	Model          string  `json:"model"`
	Year           int16   `json:"year"`
	Body           string  `json:"body"`
	Mileage        int     `json:"mileage"`
	EngineCapacity float32 `json:"engine_capacity"`
	IsAvailable    bool    `json:"isAvailable"`
	Color          string  `json:"color"`
}

func NewCar(vendor, model, body, color string, year int16, mileage int, isAvailable bool, engineCapacity float32) *Car {
	return &Car{
		Vendor:         vendor,
		Model:          model,
		Body:           body,
		Year:           year,
		IsAvailable:    isAvailable,
		Mileage:        mileage,
		EngineCapacity: engineCapacity,
		Color:          color,
	}
}

func (c *Car) Validate() error {
	if len(c.Vendor) < 1 {
		return fmt.Errorf("vendor field cannot be empty")
	}
	if len(c.Model) < 1 {
		return fmt.Errorf("model field cannot be empty")
	}
	if c.Year < 1900 && c.Year > 2024 {
		return fmt.Errorf("unappropiate year")
	}
	return nil
}
