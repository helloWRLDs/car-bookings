package domain

import (
	"net/http"
	"strconv"
	"strings"
)

type Filters struct {
	Offset     int
	Limit      int
	Sort       string
	Order      string
	Filter     string
	FilterType string
}

func ExtractFilters(r *http.Request) Filters {
	validFields := map[string]bool{
		"id":              true,
		"vendor":          true,
		"model":           true,
		"year":            true,
		"body":            true,
		"engine_capacity": true,
		"mileage":         true,
		"color":           true,
	}
	limit, _ := strconv.Atoi(r.URL.Query().Get("limit"))
	offset, _ := strconv.Atoi(r.URL.Query().Get("offset"))
	sort := r.URL.Query().Get("sort")
	order := strings.ToUpper(r.URL.Query().Get("order"))
	if limit <= 0 {
		limit = 10
	}
	if offset < 0 {
		offset = 0
	}
	if !validFields[sort] {
		sort = "id"
	}
	if order != "ASC" && order != "DESC" {
		order = "ASC"
	}
	filters := Filters{
		Limit:      limit,
		Offset:     offset,
		Sort:       sort,
		Order:      order,
		FilterType: "model",
		Filter:     "",
	}
	for key, _ := range validFields {
		if r.URL.Query().Get(key) != "" {
			filters.FilterType = key
			filters.Filter = r.URL.Query().Get(key)
		}
	}
	return filters
}
