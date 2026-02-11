package config

import (
	"flag"
	"fmt"
)


func CreateFlags() *string {
	
	environment := flag.String("environment", "dev", "the environment stage")

	flag.Parse()
	fmt.Println(*environment)

	return environment
}
