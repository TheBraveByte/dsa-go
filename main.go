package main

import (
	"fmt"

	"github.com/akinbyte/dsa-go/solutions"
)

func main() {
	fmt.Println("--- Solve Me First ---")
	var a, b, res uint32
	fmt.Scanf("%v\n%v", &a, &b)
	res = solutions.SolveMeFirst(a, b)
	fmt.Println(res)

	fmt.Println("--- Dealing with Operators ---")
	var mealCost float64
	var tipPercent, taxPercent int32
	fmt.Scanf("%v %v %v", &mealCost, &tipPercent, &taxPercent)
	solutions.Solve(mealCost, tipPercent, taxPercent)

	fmt.Println("--- Dealing with Conditional Statements ---")
	solutions.IntegerChecker()
}
