package main

import (
	"fmt"
	"net/http"

	"github.com/jjosephy/foodapi/handler"
	"github.com/jjosephy/foodapi/provider"
)

// Main entry point used to set up routes
func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/food", handler.FoodHandler(provider.NewFoodAPIProvider()))
	err := http.ListenAndServeTLS(":8443", "./cert.pem", "./private_key", mux)
	if err != nil {
		fmt.Printf("main(): %s\n", err)
	}
}
