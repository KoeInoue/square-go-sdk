package api

import (
	"github.com/KoeInoue/square-go-sdk/models"
)

// CatalogApiInterface defines the interface for the CatalogApi
type CatalogApiInterface interface {
	ListCatalog(models.ListCatalogRequest) (*models.ListCatalogResponse, error)
	RetrieveCatalogObject(objectID string) (*models.CatalogResponse, error)
}

// CatalogApi is a struct that implements CatalogApiInterface
type CatalogApi struct{}

// NewCatalogApi returns a new CatalogApi
func NewCatalogApi() *CatalogApi {
	return &CatalogApi{}
}

// CreateCatalog returns list of catalog objects
func (api *CatalogApi) ListCatalog(req models.ListCatalogRequest) (*models.ListCatalogResponse, error) {
	res := models.ListCatalogResponse{}
	return request(req, res, "/v2/catalog/list", HTTP_GET)
}

// RetrieveCatalogObject
func (api *CatalogApi) RetrieveCatalogObject(objectID string) (*models.CatalogResponse, error) {
	res := models.CatalogResponse{}
	u := "/v2/catalog/object/" + objectID

	return request(struct{}{}, res, u, HTTP_GET)
}
