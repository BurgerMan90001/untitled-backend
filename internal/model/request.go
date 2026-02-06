package model

import "time"

type Request struct {
	Format string `json:"format"` // Format, as in time.Format. If empty, use time.RFC3339.
	TZ     string `json:"tz"`     // TZ, as in time.LoadLocation. If empty, use time.Local.
}

// The time, formatted according to the request's Format and TZ.
type Resp struct {
	Time time.Time `json:"time"`
}
type Error struct {
	Error string `json:"error"`
}