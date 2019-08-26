package main

import (
	"fmt"
)

type IPAddr [4]byte

//TOD: add a "String() string"method to IPAddr
func (s IPAddr) String() string {
	return fmt.Sprintf("%v.%v.%v.%v",s[0],s[1],s[2],s[3])
}
func main() {
	addrs := map[string]IPAddr{
		"sdf":   {12, 123, 1, 12},
		"sddsf": {122, 13, 21, 122},
	}
	for a, e := range addrs {
		fmt.Printf("%v: %v\n", a, e)

	}
}
