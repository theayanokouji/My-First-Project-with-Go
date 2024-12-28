package main

import "fmt"

func main() {
	// Write your code here
	bubblegum := 202
	toffee := 118
	iceCream := 2250
	milkChocolate := 1680
	doughnut := 1075
	pancake := 80
	// calculate the total income
	income := bubblegum + toffee + iceCream + milkChocolate + doughnut + pancake
	fmt.Println("Earned amount:")
	fmt.Printf("Bubblegum: $%d\n", bubblegum)
	fmt.Printf("Toffee: $%d\n", toffee)
	fmt.Printf("Ice cream: $%d\n", iceCream)
	fmt.Printf("Milk chocolate: $%d\n", milkChocolate)
	fmt.Printf("Doughnut: $%d\n", doughnut)
	fmt.Printf("Pancake: $%d\n\n", pancake)
	fmt.Printf("Income: $%d\n", income)
	// get the staff expenses and other expenses
	var staffExpenses, otherExpenses int
	fmt.Print("Staff expenses: ")
	fmt.Scan(&staffExpenses)
	fmt.Print("Other expenses: ")
	fmt.Scan(&otherExpenses)
	// calculate net income
	income -= staffExpenses + otherExpenses
	fmt.Printf("Net income: $%d\n", income)
}
