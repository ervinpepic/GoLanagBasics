package main

import (
	"fmt"
)

const toyotaHeadQuarters, fordHeadQuarters = "Toyota Motor Corporation", "Ford Motor Corporation"

// var newHQ string = "New HQ"

func main() {
	newHQ := "Ford HQ"
	const myConst = "No Value needed"
	fmt.Println(toyotaHeadQuarters, newHQ, myConst)
}
