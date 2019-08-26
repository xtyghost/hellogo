package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m = map[string]Vertex{
	"sdfs":{123,1231},
	"sdwefs":{1,1},
}

func main()  {
	fmt.Println(m)

}