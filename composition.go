package main

import (
	"fmt"
)

func main() {
	car_a := car{}
	car_a.mpg = 25
	car_a.numberOfDors = 2
	fmt.Println(car_a)
	car_a.getMpg()

	car_b := car{vehicle{19, 4}}
	fmt.Println(car_b)
	car_b.getMpg()

}

type vehicle struct {
	mpg          int
	numberOfDors int
}

type car struct {
	vehicle
}

func (v *vehicle) getMpg() {
	println("This vehicle gets", v.mpg, " mpg")
}
