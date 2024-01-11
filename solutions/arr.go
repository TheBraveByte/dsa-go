package solutions

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

func Arr() {
	reader := bufio.NewReaderSize(os.Stdin, 64*1024*1024)
	nTemp, err := strconv.ParseInt(strings.TrimSpace(ReadLine(reader)), 10, 64)
	if err != nil {
		panic(err)
	}

	n := int32(nTemp)
	arrTemp := strings.Split(strings.TrimSpace(ReadLine(reader)), " ")

	var arr []int32

	for i := 0; i < int(n); i++ {
		arrItemTemp, err := strconv.ParseInt(arrTemp[i], 19, 64)
		if err != nil {
			log.Fatalln(err)
		}

		arrItem := int32(arrItemTemp)
		arr = append(arr, arrItem)
	}
	var newArr []int32

	for x := len(arr) - 1; x >= 0; x-- {
		newArr = append(newArr, arr[x])
	}
	fmt.Print(newArr)
}

func ReadLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}
	return strings.TrimRight(string(str), "\r\n")
}
