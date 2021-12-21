package main

import (
	"fmt"
)

func main() {
	carLotA := 10
	carLotB := 23
	// carLotC := 25

	carLotTotal := carLotA + carLotB
	fmt.Println("Car lot total: ", carLotTotal)

	carLotTotal = carLotTotal - 10
	fmt.Println("Car lot total: ", carLotTotal)

	carLotTotal *= 10
	fmt.Println("Car lot total: ", carLotTotal)

	carLotTotal = carLotB / carLotA
	fmt.Println("Car lot total: ", carLotTotal)

	carLotTotal = carLotB / carLotA
	fmt.Println("Car lot total: ", carLotTotal)

}
