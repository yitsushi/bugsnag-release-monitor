package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

type Filter struct {
	Type  string `json:"type"`
	Value string `json:"value"`
}

type FilterParameter struct {
	filters map[string][]Filter
}

func NewFilterParameter() *FilterParameter {
	return &FilterParameter{filters: map[string][]Filter{}}
}

func (f *FilterParameter) Add(key, _type, value string) {
	if _, ok := f.filters[key]; ok {
		f.filters[key] = append(f.filters[key], Filter{_type, value})
	} else {
		f.filters[key] = []Filter{
			Filter{_type, value},
		}
	}
}

func (f FilterParameter) ToURLParams() string {
	flat := make([]string, 0)
	template := "filters[%s][][%s]=%s"

	for key, values := range f.filters {
		for _, value := range values {
			flat = append(flat, fmt.Sprintf(template, key, "type", value.Type))
			flat = append(flat, fmt.Sprintf(template, key, "value", value.Value))
		}
	}

	return strings.Join(flat, "&")
}

func (f FilterParameter) ToJSON() []byte {
	data, _ := json.Marshal(f.filters)
	return data
}