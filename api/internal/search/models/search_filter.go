package models

// Filter фильтр для поиска по объектам.
type Filter struct {
	Name   string   `json:"filterName"`
	Values []string `json:"values"`
}
