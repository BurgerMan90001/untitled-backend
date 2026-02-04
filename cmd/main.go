package main

import (
	"fmt"
	"maps"

	"github.com/BurgerMan90001/untitled-backend/internal"
	test "github.com/BurgerMan90001/untitled-backend/tests"
)
const (
    www = 2
)
//const name = "writetcp"
func main() {
    internal.Goop()
    test.Request()
   //app.test()
    
	// log.SetPrefix(name + "\t")

	// // register the command-line flags: -p specifies the port to connect to
	// port := flag.Int("p", 8080, "port to connect to")
	// flag.Parse()

	// listener, err := net.DialTCP("tcp", &net.TCPAddr{Port: *port})
	// //log.Fatal("adsasd")
    // fmt.Println("Hello, world.")


	// //log.Fatal("Program Terminated")

	//test()
	//nums := []int{2,3,4}
	
	//fmt.Println(nums)
	
}	

func al() {

    m := make(map[string]int)

    m["k1"] = 7
    m["k2"] = 13

    fmt.Println("map:", m)

    v1 := m["k1"]
    fmt.Println("v1:", v1)

    v3 := m["k3"]
    fmt.Println("v3:", v3)

    fmt.Println("len:", len(m))

    delete(m, "k2")
    fmt.Println("map:", m)

    clear(m)
    fmt.Println("map:", m)

    _, prs := m["k2"]
    fmt.Println("prs:", prs)

    n := map[string]int{"foo": 1, "bar": 2}
    fmt.Println("map:", n)

    n2 := map[string]int{"foo": 1, "bar": 2}
    if maps.Equal(n, n2) {
        fmt.Println("n == n2")
    }

}
