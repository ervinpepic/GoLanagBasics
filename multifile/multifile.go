package main

import (
	"ervin.com/CarStuff"
	"fmt"
)

func main() {
	c := CarStuff.Car{4, 6}
	fmt.Println("Car struct print", c)
	fmt.Println("Number of car doors is: ", c.GetDoors())

	t := Truck{2, "full", oneTon}
	fmt.Println("Truck struct print", t)
	fmt.Println("Number of truck doors is: ", t.GetDoors())

}
