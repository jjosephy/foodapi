package handler

import (
	"net/http"

	"github.com/jjosephy/foodapi/logging"
	"github.com/jjosephy/foodapi/provider"
)

// FoodHandler provides the public API handler for all requests
func FoodHandler(provider provider.FoodDataProvider) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		d, err := provider.GetData(w, r)
		if err != nil {
			logging.WriteMessage("ERROR", err.Error())
		}

		if d == nil {
			logging.WriteMessage("ERROR", "instance is nil")
		}
	}
}
