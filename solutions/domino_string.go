package solutions

import "strings"

func DominoString(dstring string) int {
	newArr := [][]string{}

	total := 0

	count := 0
	arr := strings.Split(dstring, ",")
	for _, x := range arr {
		newArr = append(newArr, strings.Split(x, "-"))
	}
	if len(newArr) <= 1 {
		return count
	} else {
		for a, b := 0, 1; a < len(newArr)-1; a, b = a+1, b+1 {
			if newArr[a][1] == newArr[b][0] {
				total += 1
			} else {
				count = 1
				if total <= count {
					total = count
				}
			}
		}
		return total
	}
}
