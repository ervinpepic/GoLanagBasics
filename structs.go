package main

import "fmt"

func main() {
	fmt.Println("hello")
	d := dealerShip{name: "A1 Auto", city: "LosAngeles"}
	d.city = "Rozaje"
	fmt.Println("city --> " + d.city)

	var d2 dealerShip
	d2 = dealerShip{name: "Discount Auto", city: "Houston"}
	fmt.Println(d2)

	d3 := dealerShip{}
	fmt.Println("d3.name = " + d3.name)
}

type dealerShip struct {
	name string
	city string
}
