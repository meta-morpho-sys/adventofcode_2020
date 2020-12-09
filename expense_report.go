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
	expense, _ := twoSum(expensesInput, 2020)
	fmt.Println(expense)
}

// Take a list of expenses and check which two make the sum of 2020
func twoSum(expenses []int, target int) (int, int) {
    if len(expenses) <= 1 {
        return -1, -1
	}
	//	how to take an element and sum it with all the others?
	//	Check if sum equals to 2020
    m := make(map[int]int)
    for i, v := range expenses {
		m[i] = v
		fmt.Println("Map has these keys: ",m[i])
		fmt.Println(m)
		// result := target - m[i]
		// fmt.Println("Result is: ",result)

		// if result - target == 0 {
		// 	m[i] = result
		// 	fmt.Println(m)
		// }
        
    }
    return -1, -1
}