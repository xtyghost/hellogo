package main

import "fmt"

func main() {
	pow := make([]int, 10)
	for e := range pow {
		pow[e] = 1 << uint(e)
	}
	for _, value := range pow {
		fmt.Printf("%d\n", value)
	}

}
