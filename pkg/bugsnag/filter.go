package bugsnag

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Filter represents a filter value with given type.
//   Type can be: eq, ne
type Filter struct {
	Type  string `json:"type"`
	Value string `json:"value"`
}

// FilterParameter holds all the parameters together.
// As Bugsnag expect a specific format for filters and
// it was easier to make a wrapper struct to handle
// that format and manage multiple Filters for a
// single filters key.
type FilterParameter struct {
	filters map[string][]Filter
}

// NewFilterParameter is a helper function to
// create a new FilterParameter.
func NewFilterParameter() *FilterParameter {
	return &FilterParameter{filters: map[string][]Filter{}}
}

// Add simply adds a new filter by key, type and value.
// Example: Add("error.status", "eq", "open").
func (f *FilterParameter) Add(key, _type, value string) {
	if _, ok := f.filters[key]; ok {
		f.filters[key] = append(f.filters[key], Filter{_type, value})
	} else {
		f.filters[key] = []Filter{
			{_type, value},
		}
	}
}

// ToURLParams returns with a simple URL parameter representation
// of all the filter parameters.
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

// ToJSON converts the filter list to JSON. Not in use, but other
// endpoints (with POST method) may require that format.
func (f FilterParameter) ToJSON() []byte {
	data, _ := json.Marshal(f.filters)
	return data
}
