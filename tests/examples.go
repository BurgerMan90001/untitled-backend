package examples

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

func Request() {
    ctx := context.TODO()
    var body io.Reader = nil
    const method = "GET"
    const url = "https://eblog.fly.dev/index.html"

    req, _ := http.NewRequestWithContext(ctx, method, url,  body)
    fmt.Print(req)
}

func Test() {

}

type TextHandler string

func (t TextHandler) ServeHTTP(w http.ResponseWriter, _ *http.Request) { w.Write([]byte(t)) } // implicit 200 OK
func test() {
	server := http.Server{
        Addr: ":8080",
        Handler: TextHandler("hello, world"), 
        ReadTimeout: time.Minute, 
        WriteTimeout: time.Minute }
	go server.ListenAndServe()
	req, _ := http.NewRequestWithContext(context.TODO(), "GET", "http://localhost:8080", nil)
	resp, err := new(http.Client).Do(req)
	_ = err
	defer resp.Body.Close()
	resp.Write(os.Stdout) // print the response to stdout.
}