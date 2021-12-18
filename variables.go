package main

import ("fmt")

 // First way of declaring variables
	var carTypeDist1 string = "Toyota"
	var carTypeDist2 = "Ford" 


// second way of declaring variables
var carTypeDist1_2, carTypeDist2_2 = "Toyota", "Ford" 

//third way

var (
	carTypeDist1_3 = "Toyota"
	carTypeDist2_3 = "Ford"
)

func main() {
	fmt.Println(carTypeDist2, carTypeDist1)
}