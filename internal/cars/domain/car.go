package domain

type Car struct {
	ID          int64  `json:"id,omitempty"`
	Vendor      string `json:"vendor"`
	Model       string `json:"model"`
	Year        int16  `json:"year"`
	IsAvailable bool   `json:"isAvailable"`
	Body        string `json:"body"`
}

func NewCar(vendor, model, body string, year int16, isAvailable bool) *Car {
	return &Car{
		Vendor:      vendor,
		Model:       model,
		Body:        body,
		Year:        year,
		IsAvailable: isAvailable,
	}
}
