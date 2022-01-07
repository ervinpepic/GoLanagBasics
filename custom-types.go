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

	car_b := car{vehicle{19, 4}, red}
	fmt.Println(car_b)
	car_b.getMpg()
	car_b.color = red
	car_c := car{}
	car_c.color = black
	car_c.mpg = 34
	car_c.getMpg()
	fmt.Println(car_c)

}

type vehicle struct {
	mpg          int
	numberOfDors int
}

type car struct {
	vehicle
	color Color
}

func (v *vehicle) getMpg() {
	println("This vehicle gets", v.mpg, " mpg")
}

//custom-types

type Color string

const (
	blue  Color = "blue"
	red   Color = "red"
	black Color = "black"
)
