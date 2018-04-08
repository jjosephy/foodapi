package provider

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/jjosephy/foodapi/errorcodes"
	"github.com/jjosephy/foodapi/logging"
)

// Search API template
const searchTemplate string = "https://api.nal.usda.gov/ndb/search/?format=json&q=%s&sort=n&max=25&offset=0&api_key=DEMO_KEY"

// FoodDataProvider is the public interface for getting food data
type FoodDataProvider interface {
	GetData(w http.ResponseWriter, r *http.Request) (map[string]interface{}, error)
}

// FoodAPIProvider implementation of FoodDataProviderInterface
type FoodAPIProvider struct {
}

// NewFoodAPIProvider returns an instance of FoodAPIProvider that implements FoodDataProvider interface
func NewFoodAPIProvider() *FoodAPIProvider {
	return &FoodAPIProvider{}
}

// TODO: Create Common Response Model with versions

// GetData public API implementation
func (f *FoodAPIProvider) GetData(w http.ResponseWriter, r *http.Request) (map[string]interface{}, error) {

	logging.WriteMessage("INFO", "started")
	var data map[string]interface{}
	var err error
	var body []byte

	query := r.URL.Query().Get("s")
	if query == "" {
		return nil, handleError(http.StatusBadRequest, w, errorcodes.ErrorInvalidSearchQuery)
	}

	// Switch on Request Method
	switch r.Method {
	case "GET":
		s := fmt.Sprintf(searchTemplate, query)
		resp, err := http.Get(s)
		if err != nil {
			return nil, handleError(http.StatusInternalServerError, w, errorcodes.ErrorFailedCallingAPI)
		}

		defer resp.Body.Close()
		body, err = ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, handleError(http.StatusBadRequest, w, errorcodes.ErrorCouldNotReadResponseBody)
		}

		if err := json.Unmarshal(body, &data); err != nil {
			return nil, handleError(http.StatusBadRequest, w, errorcodes.ErrorCouldNotParseResponseBody)
			w.WriteHeader(http.StatusBadRequest)
		}

		w.Write(body)
		return data, nil
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		return data, err
	}
}

// Private method to handle errors
func handleError(statusCode int, writer http.ResponseWriter, errorMsg string) error {
	writer.WriteHeader(statusCode)
	fmt.Fprintf(writer, errorMsg)
	return errors.New("invalid search query")
}
