package CarStuff

type Car struct {
	NumberOfDors int
	Cylinders    int
}

func (c *Car) GetDoors() int {
	return c.NumberOfDors
}
