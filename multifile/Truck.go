package main

type Truck struct {
	numberOfDors int
	bedSize      string
	weightByTon  weight
}

type weight string

const (
	oneTon weight = "One Ton"
	TwoTon weight = "Two Ton"
)

func (t *Truck) GetDoors() int {
	return t.numberOfDors
}
