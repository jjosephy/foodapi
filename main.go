package main

import (
	"fmt"
	"net/http"

	"github.com/jjosephy/foodapi/handler"
)

// Main entry point used to set up routes
func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/food", handler.FoodHandler())
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		fmt.Printf("main(): %s\n", err)
	}
}
