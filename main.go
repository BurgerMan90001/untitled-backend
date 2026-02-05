package main

import (
	"flag"
	"fmt"


)

func readEnvFile() {
	path := filepath.Join(os.TempDir(), "dat")
    dat, err := os.ReadFile(path)
    checkErr(err)
    fmt.Print(string(dat))
}
func main() {
    test := flag.String("timeout", "asdasd", "timeout for connecting to postgres")
    flag.Parse()
    fmt.Println(*test)
    server.Run()
}
