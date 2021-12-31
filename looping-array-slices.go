package main

import (
	"fmt"
)

func main() {
	var carTypes [3]string
	carTypes[0] = "Passat"
	carTypes[1] = "E class"
	carTypes[2] = "M200"

	i := 0
	for i <= len(carTypes)-1 {
		fmt.Println("Car types are: ", carTypes[i])
		i++
	}
	//also
	for i < len(carTypes) {
		fmt.Println("Car types are: ", carTypes[i])
		i++
	}

	fmt.Println("\n For Loop Verobse")

	for k := 0; k < len(carTypes); k++ {
		fmt.Println("Cars: ", carTypes[k])
	}

	fmt.Println("\n For Loop Range")

	for key, value := range carTypes {
		fmt.Println("Key is: ", key, "Value is: ", value)
	}

	fmt.Println("\n For Loop Range ignore key-value")

	for _, key := range carTypes {
		fmt.Println("Key is: ", "Value is: ", key)
	}
}
