package responses

// The time, formatted according to the request's Format and TZ.
type Response struct {
	Message string `json:"message"`
	Status string `json:"status"`
}
type Error struct {
	Error string `json:"error"`
	Status int `json:"status"`  
}

type DeleteResponse struct {
	Message string `json:"message"`
	Status string `json:"status"`
}

