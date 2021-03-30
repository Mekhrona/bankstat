package stats

import (
	"fmt"

	"github.com/Mekhrona/bankstat/pkg/bank/types")


func ExampleAvg() {

	payments := [] types.Payment {

		{
			ID: 56,
			Amount: 55_00,
			Category: "food",
		},
		{
			ID: 79,
			Amount: 165_00,
			Category: "clothes",
		},
	}
	
	fmt.Println(Avg(payments))
	//Output:11000

}