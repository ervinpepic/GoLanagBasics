package main

import (
	"fmt"
)

func main() {
	person_instace := new(person) //here we are creating a person instance (struct person)
	fmt.Println(person_instace.talk())
	fmt.Println(person_instace.walk())

	police_officer_instace := new(policeOfficer) //here we are creating an office instance (struct person)
	fmt.Println(police_officer_instace.talk())
	fmt.Println(police_officer_instace.walk())
	fmt.Println(police_officer_instace.run())
	police_officer_instace.bagdeNumber = 150
}

type person struct {
	firstname string
	lastname  string
}

type policeOfficer struct {
	bagdeNumber int
	precint     string
}

type behaviors interface {
	walk() int
	talk() string
	run() int
}

//person struct implementation
func (person_func_argumet *person) talk() string { //this function takes PERSON STRUCT as function argument and implenets talk method from the interface
	return "Hello there!"
}

func (person_func_argumet *person) walk() int { //this function takes PERSON STRUCT as function argument and implenets walk method from the interface
	return 10
}

//police officer implementation

func (police_ofc *policeOfficer) talk() string { //this function takes policeOfficer STRUCT as function argument and implenets talk method from the interface
	return "Mey I have your license, please?"
}

func (police_ofc *policeOfficer) walk() int { //this function takes policeOfficer STRUCT as function argument and implenets talk method from the interface
	return 120
}
func (police_ofc *policeOfficer) run() int { //this function takes policeOfficer STRUCT as function argument and implenets talk method from the interface
	return 1700
}
