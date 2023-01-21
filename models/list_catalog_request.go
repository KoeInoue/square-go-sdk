package models

type ListCatalogRequest struct {
	cursor         string
	types          string
	catalogVersion int64
}
