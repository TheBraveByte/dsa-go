package solutions

import (
	"fmt"
)

func Bin() {
	var zeroCount int
	var oneCount int

	mapBin := make(map[string]int)
	binValue := fmt.Sprintf("%b", 125)

	fmt.Println(binValue)

	for _, y := range binValue {
		mapBin[string(y)]++
	}
	oneCount, ok := mapBin["1"]
	zeroCount, ok = mapBin["0"]

	if !ok {
		fmt.Println("Invalid Binary input :: check input")
	}

	fmt.Println(oneCount - zeroCount)
	fmt.Println(mapBin)
}

func NewBin() {
	count := 0
	var oneArr []int

	baseTwo := fmt.Sprintf("%b", 157383)

	fmt.Println(baseTwo)

	for _, x := range baseTwo {

		if x == '1' {
			count += 1
		} else {
			count = 0
		}
		oneArr = append(oneArr, count)
	}

	for _, y := range oneArr {
		if y >= count {
			fmt.Println(y)
		}
	}
}
