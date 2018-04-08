package handler

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/jjosephy/foodapi/test"
)

const (
	uriTemplate = "Json"
)

// Test Handles
var h http.HandlerFunc
var ts *httptest.Server

func TestMain(m *testing.M) {
	h = FoodHandler(test.NewTestFoodAPIProvider())
	ts = httptest.NewServer(http.HandlerFunc(h))
	defer ts.Close()
	os.Exit(m.Run())
}
