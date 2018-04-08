package handler

import (
	"net/http"

	"github.com/jjosephy/foodapi/provider"
)

// FoodHandler provides the public API handler for all requests
func FoodHandler(provider provider.FoodDataProvider) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		provider.GetData(w, r)
	}
}
