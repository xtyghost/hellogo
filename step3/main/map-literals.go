package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m = map[string]Vertex{
	"Bell Labs": {123, 123},
	"Google":    {213, 1231},
}

func main() {
	fmt.Println(m)
	for e := range m {
		fmt.Printf("%s=================",e,m[e])
		println()
	}

}
