package main

import (
	"fmt"
)

var shippingDays = 30
var shippingDaysPointer = new(int)

func main() {
	shippingDaysAdjustments(shippingDays)
	fmt.Println("after shippingDaysAdjustments function call", shippingDays)

	shippingDaysAdjustmentsPointer(&shippingDays)
	fmt.Println("after shippingDaysAdjustments pointer function call", shippingDays)
	fmt.Println("after shippingDaysAdjustments pointer function call reference", &shippingDays)

	shipper := shipper{}
	shipper.id = 320
	shipper.name = "Pacific Shipper"
	shipperUpdates(&shipper)
	fmt.Println("shipper.id after ShipperUpdates function call -> ", shipper.id)
	fmt.Println("shipper.name after ShipperUpdates function call -> ", shipper.name)

	*shippingDaysPointer = 55
	shippingDaysAdjustmentsPointer(shippingDaysPointer)
	fmt.Println("ShippingDaysPointer ShippingDaysAdjustments pointer function call", *shippingDaysPointer)
}

func shippingDaysAdjustments(shippingDays int) {
	shippingDays += 10
}

func shippingDaysAdjustmentsPointer(shippingDays *int) {
	*shippingDays += 10
}

func shipperUpdates(s *shipper) {
	s.id += 10
	s.name += "Inc."
}

type shipper struct {
	name string
	id   int
}
