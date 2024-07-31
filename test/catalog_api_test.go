package test

import (
	"testing"
)

func TestCatalogApi(t *testing.T) {
	testPlanId := "HLXWM2ZDM4ERD5PRTSYBOKNC"
	cRes, err := client.CatalogApi.RetrieveCatalogObject(testPlanId)
	if err != nil {
		t.Error(err)
	}

	if len(cRes.Errors) > 0 {
		t.Errorf("Square Error: %#v, Code: %s", cRes.Errors[0].Detail, cRes.Errors[0].Code)
	}
}
