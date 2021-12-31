package main

import (
	"fmt"
)

func main() {
	carDealerInventory := make(map[string]int)
	carDealerInventory["A1 Auto"] = 76
	fmt.Println("A1 inventory value is: ", carDealerInventory["A1 Auto"])
	fmt.Println("Care Dealr inventory lenght: ", len(carDealerInventory))

	carDealerInventory["A2 Auto"] = 111
	fmt.Println("A2 inventory value is: ", carDealerInventory["A2 Auto"])
	fmt.Println("Care Dealr inventory lenght: ", len(carDealerInventory))
	delete(carDealerInventory, "A1 Auto")
	fmt.Println("Care Dealr inventory lenght after delete: ", len(carDealerInventory))
	fmt.Println("A1 auto value: ", carDealerInventory["A1 Auto"])
	fmt.Println("A2 auto value: ", carDealerInventory["A2 Auto"])

	carDealerTown := make(map[string]string)
	carDealerTown["A1 Auto"] = "Austin Texas"
	carDealerTown["A2 Auto"] = "Rozaje Grahovo"
	fmt.Println("A1 delaer auto town is: ", carDealerTown["A1 Auto"])

}
