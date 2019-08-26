package main

import "fmt"

func main()  {
	p :=[]int{1,24,5,6,8,9,0}
	fmt.Printf("p=========%d",p)
	fmt.Println()
	fmt.Println("p[:3]==",p[:3])
	fmt.Println("p[4:]==",p[4:])
	fmt.Println("p[0:len(p)]==",p[0:])

}
