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
	expense, err := ShowExpenses(expensesInput)
	if err != nil {
		fmt.Printf("ERROR: %s", err)
		return
	 }

	fmt.Println(expense)
}

func ShowExpenses(expenses []int) (expense int, err error) {
	sum := 0
	for _, expense := range expenses {
		fmt.Println(expense)		
		sum += expense
	}
	return sum, nil
}
