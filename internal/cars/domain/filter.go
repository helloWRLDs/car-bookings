package domain

import (
	"fmt"
	"net/http"
	"net/url"
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

func ExtractFilters(r *http.Request) *Filters {
	var (
		order      = strings.ToUpper(r.URL.Query().Get("order"))
		limit, _   = strconv.Atoi(r.URL.Query().Get("limit"))
		offset, _  = strconv.Atoi(r.URL.Query().Get("offset"))
		sort       = r.URL.Query().Get("sort")
		filter     = "'%%'"
		filterType = "model"
	)
	validateFilters(r.URL.Query(), &limit, &offset, &sort, &order, &filter, &filterType)
	return &Filters{
		Limit:      limit,
		Offset:     offset,
		Sort:       sort,
		Order:      order,
		Filter:     filter,
		FilterType: filterType,
	}
}

// set default values if missing filters
func validateFilters(url url.Values, limit, offset *int, sort, order, filter, filterType *string) {
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

	if *limit <= 0 {
		*limit = 10
	}
	if *offset < 0 {
		*offset = 0
	}
	if *sort == "" {
		*sort = "id"
	}
	if *order != "ASC" && *order != "DESC" {
		*order = "ASC"
	}
	for key, _ := range validFields {
		if url.Get(key) != "" && len(url.Get(key)) < 10 {
			*filterType = fmt.Sprintf("%s::text", key)
			*filter = "'%" + url.Get(key) + "%'"
		}
	}
}
