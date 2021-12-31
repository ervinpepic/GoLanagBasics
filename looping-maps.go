package main

import (
	"fmt"
)

func main() {

	carDealerInventory := make(map[string]int)
	carDealerInventory["A1 Auto"] = 76
	carDealerInventory["A2 Auto"] = 126

	carDealerInventory2 := make(map[int]string)
	carDealerInventory2[76] = "A1 Auto"
	carDealerInventory2[112] = "A2 Auto"

	// for i := 0; i < len(carDealerInventory2); i++ {
	// 	fmt.Println(carDealerInventory2[i])
	// } WE CANT LOOP THROGUH MAPS LIKE THIS

	for i := range carDealerInventory {
		fmt.Println(i, " -> ", "Storage supplies: ", carDealerInventory[i])
	}

}
