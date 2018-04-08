package errorcodes

// Error codes for public operations
const (
	ErrorNoQuery = 1000
)

// Error messages for public operations
const (
	ErrorInvalidSearchQuery        = "invalid search query"
	ErrorCouldNotReadResponseBody  = "could not read response body"
	ErrorCouldNotParseResponseBody = "could not parse the request body"
	ErrorFailedCallingAPI          = "failed calling api"
)
