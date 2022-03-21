package controllers

// response : Struct for API response.
type response struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Result  interface{} `json:"result"`
}

// newResponse : Return new response.
func newResponse(status int, message string, result interface{}) *response {
	return &response{status, message, result}
}
