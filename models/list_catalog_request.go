package models

type ListCatalogRequest struct {
	Cursor         string `json:"cursor,omitempty"`
	Types          string `json:"types,omitempty"`
	CatalogVersion int64  `json:"catalog_version,omitempty"`
}
