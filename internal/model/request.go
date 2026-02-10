package model

type Request struct {
	Format string `json:"format"` // Format, as in time.Format. If empty, use time.RFC3339.
	TZ     string `json:"tz"`     // TZ, as in time.LoadLocation. If empty, use time.Local.
}

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

