// Source: HackerRank
// Synopsis: A function that returns the sum of two integers.

/*
Complete the function solveMeFirst to compute the sum of two integers.

# Example

Return .

# Function Description

Complete the solveMeFirst function in the editor below.

solveMeFirst has the following parameters:

int a: the first value
int b: the second value
Returns
- int: the sum of  and

# Constraints

# Sample Input

a = 2
b = 3
Sample Output

5
*/
package solutions

import "fmt"

func SolveMeFirst(a uint32, b uint32) uint32 {
	fmt.Printf("The Sum of %v and %v is equal to ", a, b)
	return a + b
}
