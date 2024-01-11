package solutions

func add(a []int) int {
	addValue := 0
	for _, v := range a {
		addValue += v
	}
	return addValue
}

func MaxIntChecker(intSlice []int) int {
	var max int
	for _, y := range intSlice {
		if y > max {
			max = y
		}
	}
	return max
}

func CountMaxHourGlass(hgArr [][]int) int {
	// valMap := make(map[int]int)

	var storeArr [][]int
	var hgValue []int

	arrRow := [][]int{{0, 1, 2}, {1, 2, 3}, {2, 3, 4}, {3, 4, 5}}
	for _, v := range arrRow {
		for x, y := range v {
			for j := 0; j < 4; j++ {
				if x == 1 {
					storeArr = append(storeArr, hgArr[y][j+1:j+2])
				} else {
					storeArr = append(storeArr, hgArr[y][j:j+3])
				}
			}
		}
	}

	for i := 0; i < len(storeArr); i++ {
		j, k := 0, 0
		if i%12 == 0 {
			j = i + 4
			k = i + 8
			for x := 0; x < len(arrRow); x++ {
				hgSum := add(storeArr[i+x]) + add(storeArr[j+x]) + add(storeArr[k+x])
				hgValue = append(hgValue, hgSum)
			}

		}
	}

	// fmt.Println("The sum of each hourglass",hgValue)
	return MaxIntChecker(hgValue)
}
