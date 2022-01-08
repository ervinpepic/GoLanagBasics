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

	car_default := newCar()
	fmt.Println("Car with newFunction default values", car_default)

	//polymorphic
	truck_a := truck{vehicle{15, 2}, black}
	mans := [...]manufacturer{car_c, truck_a}
	for i := range mans {
		fmt.Println("interaction = ", i, mans[i].getVIN())
	}

}

type vehicle struct {
	mpg          int
	numberOfDors int
}

type car struct {
	vehicle
	color Color
}

type truck struct {
	vehicle
	color Color
}

func (v *vehicle) getMpg() {
	println("This vehicle gets", v.mpg, " mpg")
}

func (t truck) getVIN() string {
	return "truck VIN"
}
func (c car) getVIN() string {
	return "car VIN"
}

type manufacturer interface {
	getVIN() string
}

//custom-types

type Color string

const (
	blue  Color = "blue"
	red   Color = "red"
	black Color = "black"
)

//default value initialization (this function let us create deafulte values for our CAR struct,
//we just need to call this function) no need to initialize the whole new struct
func newCar() *car {
	result := car{}
	result.mpg = 20
	result.numberOfDors = 4
	result.color = black
	return &result
}
