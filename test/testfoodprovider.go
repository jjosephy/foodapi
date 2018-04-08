package test

import (
	"fmt"
	"net/http"
)

// TestFoodAPIProvider implementation of FoodDataProviderInterface
type TestFoodAPIProvider struct {
}

// NewTestFoodAPIProvider returns an instance of FoodAPIProvider that implements FoodDataProvider interface
func NewTestFoodAPIProvider() *TestFoodAPIProvider {
	return &TestFoodAPIProvider{}
}

// GetData public API implementation
func (f *TestFoodAPIProvider) GetData(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "test successful")
}
