package main

import (
	"fmt"
)

func main() {
	const (
		Sunday int = iota
		Monday
		Tuesday
		Wednesday
		Thursday
		Friday
		Saturday
	)

	for i := 0; i < 7; i++ {
		switch i {
		case Sunday:
			fmt.Println("Sunday")
		case Monday:
			fmt.Println("Monday")
		case Tuesday:
			fmt.Println("Tuesday")
		case Wednesday:
			fallthrough
		case Thursday:
			fmt.Println("Thursday")
		case Friday:
			fmt.Println("Friday")
		case Saturday:
			fmt.Println("Saturday")
		}

	}

}
