package main

import "fmt"

var expensesInput = []int{2, 3}

func main() {
	expense := ShowExpenses(expensesInput)
	fmt.Println(expense)
}

func ShowExpenses(expenses []int) (expense int) {
	sum := 0
	for _, expense := range expenses {
		sum += expense
	}
	return sum
}
