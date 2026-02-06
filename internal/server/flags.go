package server

import (
	"flag"
	"fmt"
)


func test123() {
	test := flag.String("timeout", "asdasd", "timeout for connecting to postgres")
	flag.Parse()
	fmt.Println(*test)
}
