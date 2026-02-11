package model


type Request struct {
	Format string `json:"format"` // Format, as in time.Format. If empty, use time.RFC3339.
	TZ     string `json:"tz"`     // TZ, as in time.LoadLocation. If empty, use time.Local.
}