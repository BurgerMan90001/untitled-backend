package config

import (
	"flag"
	"fmt"
)


func CreateFlags() {
	test := flag.String("timeout", "asdasd", "timeout for connecting to postgres")
	flag.Parse()
	fmt.Println(*test)
}
