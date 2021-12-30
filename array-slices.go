package main

import (
	"fmt"
)

func main() {
	//First Way of initializing array
	var carTypes [3]string

	carTypes[0] = "VW"
	carTypes[1] = "Ford"
	carTypes[2] = "Toyota"

	fmt.Println(carTypes[1])

	//second way of initializing array

	carTypes2 := [3]string{"VW", "Ford", "Toyota"}
	fmt.Println(carTypes2[0])

	//ArraySlices
	carTypesSlices := []string{"PONTIAC", "FERRARI", "LAMBORGHINI"}
	fmt.Println(carTypesSlices[2])
	fmt.Println("Before Append", len(carTypesSlices))

	carTypesSlices = append(carTypesSlices, "Tesla")
	fmt.Println("After Append", len(carTypesSlices))

	carTypeSlicesMake := make([]string, 3)
	carTypeSlicesMake[0] = "VW"
	carTypeSlicesMake[1] = "Ford"
	carTypeSlicesMake[2] = "Toyota"
	fmt.Println("Before slicing", carTypeSlicesMake)

	carTypesSlices2 := carTypesSlices[2:4]
	fmt.Println("CarTypesSlices2 length is: ", len(carTypesSlices2))
	fmt.Println("Slice[2:4] = ", carTypesSlices2)
}
