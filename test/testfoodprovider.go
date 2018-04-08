package test

import (
	"fmt"
	"net/http"
)

// FoodTestProvider implementation of FoodDataProviderInterface
type FoodTestProvider struct {
}

// NewTestFoodAPIProvider returns an instance of FoodAPIProvider that implements FoodDataProvider interface
func NewTestFoodAPIProvider() *FoodTestProvider {
	return &FoodTestProvider{}
}

// GetData public API implementation
func (f *FoodTestProvider) GetData(w http.ResponseWriter, r *http.Request) (map[string]interface{}, error) {
	fmt.Fprintf(w, "test successful")
	m := make(map[string]interface{})
	return m, nil
}
