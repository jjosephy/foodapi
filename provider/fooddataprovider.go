package provider

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// Search API template
const searchTemplate string = "https://api.nal.usda.gov/ndb/search/?format=json&q=%s&sort=n&max=25&offset=0&api_key=DEMO_KEY"

// FoodDataProvider is the public interface for getting food data
type FoodDataProvider interface {
	GetData(w http.ResponseWriter, r *http.Request)
}

// FoodAPIProvider implementation of FoodDataProviderInterface
type FoodAPIProvider struct {
}

// NewFoodAPIProvider returns an instance of FoodAPIProvider that implements FoodDataProvider interface
func NewFoodAPIProvider() *FoodAPIProvider {
	return &FoodAPIProvider{}
}

// GetData public API implementation
func (f *FoodAPIProvider) GetData(w http.ResponseWriter, r *http.Request) {

	query := r.URL.Query().Get("s")
	if query == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "invalid search query")
	}

	// Switch on Request Method
	switch r.Method {
	case "GET":
		s := fmt.Sprintf(searchTemplate, query)
		resp, err := http.Get(s)
		if err != nil {
			// handle error
			fmt.Printf("%s\n", "error")
		}
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		w.Write(body)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
}
