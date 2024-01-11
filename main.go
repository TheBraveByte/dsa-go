/*
Complete the function solveMeFirst to compute the sum of two integers.

Example


Return .

Function Description

Complete the solveMeFirst function in the editor below.

solveMeFirst has the following parameters:

int a: the first value
int b: the second value
Returns
- int: the sum of  and

Constraints


Sample Input

a = 2
b = 3
Sample Output

5
*/

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
