package handler

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// FoodHandler provides the public API handler for all requests
func FoodHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		query := r.URL.Query().Get("s")
		if query == "" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		// Switch on Request Method
		switch r.Method {
		case "GET":
			s := fmt.Sprintf("https://api.nal.usda.gov/ndb/search/?format=json&q=%s&sort=n&max=25&offset=0&api_key=DEMO_KEY", query)
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
}
