package main

import "fmt"

func main()  {
	var a []int
	printSlice("a",a)

	//the slice grows as needed
	a =append(a,1)
	printSlice("a",a)
	a=append(a,3,6,7,7,7,9)
	printSlice("a",a)
}
func printSlice(s string, x []int) {
	fmt.Printf("%s len%d cap=%d %v\n",s,len(x),cap(x),x)
}
