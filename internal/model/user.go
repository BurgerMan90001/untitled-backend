package model

import "time"


type User struct {
    Id string `json:"id"`
    Firstname string `json:"firstname"`
    Lastname  string `json:"lastname"`
    Age       int    `json:"age"`
}


/* Create a user example
var user = User {
    Id: "123",
    Firstname: "dasdasd",
    Lastname: "Wwww",
    Age: 32,
}
*/

type Request struct {
	Format string `json:"format"` // Format, as in time.Format. If empty, use time.RFC3339.
	TZ     string `json:"tz"`     // TZ, as in time.LoadLocation. If empty, use time.Local.
}

// The time, formatted according to the request's Format and TZ.
type Resp struct {
	Time time.Time `json:"time"`
} // no need for omitempty here; we'll never send a zero time.
type Error struct {
	Error string `json:"error"`
} // no need for omitempty here; we'll never send an empty error.
// Responses are written to the HTTP response body in JSON format.