package main

import "fmt"

func main()  {
	var z []int
	fmt.Println(z,len(z),cap(z))
	if z==nil {
		fmt.Printf("nil!")
	}

}
