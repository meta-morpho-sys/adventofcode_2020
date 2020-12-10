//  Day 1
//	Find the two entries that sum to 2020;
//	what do you get if you multiply them together?

package main

import (
	"fmt"
)

var expensesInput = []int{
	1721,
	979,
	366,
	299,
	675,
	1456,
}

func main() {
	expense := find2020(expensesInput)
	fmt.Println(expense)
}

// Take a list of expenses and check which two make the sum of 2020
func find2020(expenses []int) (found[]int) {
	// if len(expenses) <= 1 {
	// 	return -1, nil
	// }
	// n = 0 // Where n is the index where the value that we need to sum up is stored
	// result = xn + x(n+1)
	//	Take the n0 and sum it up with the values of indices "n" (by iterating over them?)
	//	Then pass on the n1, and so on.
	//	Nested loops?
	for i := 0; i < len(expenses); i++ {
		for i := 0; i < len(expenses); i++ {
			result := expenses[0] + expenses[i]
			fmt.Println("\ni is :", expenses[0], "next i is: ", expenses[i])
			fmt.Println("Result is: ", result)
			found := []int{}
			if result == 2020 {
				found = append(found, expenses[0], expenses[i])
				fmt.Printf("These values %v sum up to 2020",found)
			}
		}
	}
	return found
}
