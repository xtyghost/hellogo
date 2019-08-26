package main

import "fmt"

func fibonacci() func(int) int {

	return func(x int) int {
		if x < 2 {
			return 1;
		}
		t := fibonacci();
		return t(x-1) + t(x-2);
	};
}
func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f(i))
	}

}
