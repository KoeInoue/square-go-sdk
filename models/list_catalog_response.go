package models

type ListCatalogResponse struct {
	errors  []Error
	cursor  string
	objects []CatalogObject
}
