package main

import (
	"fmt"
	"os"
	// "time"
)

func main() {
	cars := fillCars()

	go showCars(cars, "first goroutine")
	go showCars(cars, "second goroutine")
	showCars(cars, "no goroutine")

	go func(msg string) {
		fmt.Println(msg)
	}("going")

	// time.Sleep(2 * time.Second)

	var input string
	fmt.Scanln(&input)
	fmt.Println("Done")
}

type Cars map[string]int

func fillCars() map[string]int{
	cars := make(map [string]int)
	cars["Jeep"] = 23
	cars["VW"] = 11
	cars["Audi"] = 50
	cars["Chevrolet"] = 10
	cars["Nissan"] = 26

	return cars
}

func showCars(c Cars, m string) {
	for key, i := range c {
		fmt.Fprintf(os.Stdout, "[%[1]v] %[2]v = %[3]v %[4]v\n", m, i, key, c[key])
	}
}