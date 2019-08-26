package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

func main() {
	a := Person{"sdf", 234}
	z := Person{"dsfs", 2342}
	fmt.Println(a, z)

}
