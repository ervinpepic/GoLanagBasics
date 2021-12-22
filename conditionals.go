package main

import (
	"fmt"
)

func main() {
	carLotA := 10
	carLotB := 20

	if carLotA < carLotB {
		fmt.Println("carLotB is greater than A")
	} else {
		fmt.Println("CarLotA is Greater")
	}

	if carLotA > 15 {
		fmt.Println("Value is greater than 15")
	} else if carLotA > 9 {
		fmt.Println("Value is greater than 9")
	} else {
		fmt.Println("Default...", carLotA)
	}

	switch carLotA {
	case 15:
		fmt.Println("Case value 15")
	case 9, 11, 12:
		fmt.Println("Case is 9 or 11 or 12")
	default:
		fmt.Println("Default case--")
		return
	}
}
