package contract

// ErrorContract for public API
type ErrorContract struct {
	Message string `json:"message"`
	Code    int    `json:"errorCode"`
}
