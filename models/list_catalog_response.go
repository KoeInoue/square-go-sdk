package models

type ListCatalogResponse struct {
	Errors  []Error         `json:"errors,omitempty"`
	Cursor  string          `json:"cursor,omitempty"`
	Objects []CatalogObject `json:"objects,omitempty"`
}
