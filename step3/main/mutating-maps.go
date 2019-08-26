package main

import "fmt"

func main() {
	m := make(map[string]int)
	m["Answer"] = 32
	fmt.Println("the value:", m["Answer"])

	m["Answer"] = 33
	fmt.Println("the value:", m["Answer"])
	delete(m, "Answer")
	fmt.Println("the value:", m["Answer"])
    v,ok :=m["Answer"]
    fmt.Println("The value:",v,"Present?",ok)
}
