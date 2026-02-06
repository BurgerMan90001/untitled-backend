package tests

import (
	"bytes"
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/BurgerMan90001/untitled-backend/internal/model"
)

type Response struct{ Time string } // note: no JSON tags here, so we just use 'Time' in uppercase.

// http handler: writes current time as JSON object (`{"Time": <time>}`)
func getTime(w http.ResponseWriter, r *http.Request) {
	var req model.Request
	w.Header().Set("Content-Type", "encoding/json")
	// if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
	// 	w.WriteHeader(400) // bad request
	// 	json.NewEncoder(w).Encode(Error{err.Error()})
	// 	return
	// }
	r.Body.Close()
	var tz *time.Location = time.Local
	if req.TZ != "" {
		var err error
		tz, err = time.LoadLocation(req.TZ)
		if err != nil || tz == nil {
			w.WriteHeader(400) // bad request
			json.NewEncoder(w).Encode(model.Error{err.Error()})
			return
		}
	}
	format := time.RFC3339
	if req.Format != "" {
		format = req.Format
	}

	resp := Response{time.Now().In(tz).Format(format)}
	json.NewEncoder(w).Encode(resp)

}

var client = &http.Client{Timeout: 2 * time.Second}

func sendRequest(tz, format string) {
	body := new(bytes.Buffer)
	json.NewEncoder(body).Encode(model.Request{TZ: tz, Format: format})
	log.Printf("request body: %v", body)
	req, err := http.NewRequestWithContext(context.TODO(), "GET", "http://localhost:8080", body)
	if err != nil {
		panic(err)
	}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	resp.Write(os.Stdout)
	resp.Body.Close()
}