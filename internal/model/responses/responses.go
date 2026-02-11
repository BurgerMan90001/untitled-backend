package responses

// The time, formatted according to the request's Format and TZ.
type Response struct {
	Message string `json:"message"`
}
type Error struct {
	Error string `json:"error"`
}

type DeleteResponse struct {
	Message string `json:"message"`
}
