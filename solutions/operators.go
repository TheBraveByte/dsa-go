/*
In this challenge, you will work with arithmetic operators. Check out the Tutorial tab for learning materials and an instructional video.

Task
Given the meal price (base cost of a meal), tip percent (the percentage of the meal price being added as tip), and tax percent (the percentage of the meal price being added as tax) for a meal, find and print the meal's total cost. Round the result to the nearest integer.

Example
meal cost = 100
tip percent = 15
tax percent = 8


A tip of 15% * 100 = 15, and the taxes are 8% * 100 = 8. Print the value  and return from the function.

Function Description
Complete the solve function in the editor below.

solve has the following parameters:

int meal_cost: the cost of food before tip and tax
int tip_percent: the tip percentage
int tax_percent: the tax percentage
Returns The function returns nothing. Print the calculated value, rounded to the nearest integer.

Note: Be sure to use precise values for your calculations, or you may end up with an incorrectly rounded result.

Input Format

There are  lines of numeric input:
The first line has a double, Meal cost (the cost of the meal before tax and tip).
The second line has an integer, Tip Percent(the percentage of  being added as tip).
The third line has an integer,  Tax Percent (the percentage of  being added as tax).

Sample Input
12.00
20
8

*/

/*
 * Complete the 'solve' function below.
 *
 * The function accepts following parameters:
 *  1. DOUBLE meal_cost
 *  2. INTEGER tip_percent
 *  3. INTEGER tax_percent
 */
package solutions

import (
	"fmt"
	"math"
)

func Solve(mealCost float64, tipPercent int32, taxPercent int32) {
	// Write your code here
	tip := mealCost * (float64(tipPercent) / 100)
	tax := mealCost * (float64(taxPercent) / 100)

	finalMealCost := int32(math.Round(tip + tax + mealCost))
	fmt.Println(finalMealCost)
}
