package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("when's Saturday")
	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("tomorrow.")
	case today + 2:
		fmt.Println("in two days.")
	default:
		fmt.Println("too far away.")

	}

}
