package examples

import (
	"context"
	"fmt"
	"io"
	"net/http"
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

func test() {
    
}