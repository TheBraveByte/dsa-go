/*

In this challenge, we learn about conditional statements. Check out the Tutorial tab for learning materials and an instructional video.

Task
Given an integer, , perform the following conditional actions:

If n is odd, print Weird
If n is even and in the inclusive range of  to , print Not Weird
If n is even and in the inclusive range of  to , print Weird
If n is even and greater than , print Not Weird
Complete the stub code provided in your editor to print whether or not  is weird.

Input Format

A single line containing a positive integer, .

Constraints
1 <= n <= 100
Output Format

Print Weird if the number is weird; otherwise, print Not Weird.

Sample Input 0

3
Sample Output 0

Weird
Sample Input 1

24
Sample Output 1

Not Weird
Explanation
`:w http.ResponseWriter, r *http.Request`
Sample Case 0:  n = 3
 n is odd and odd numbers are weird, so we print Weird.

Sample Case 1: n = 24
 and  is even, so it is not weird. Thus, we print Not Weird.
*/

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

func IntegerChecker() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)
	NTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	if err != nil {
		log.Fatal(err)
	}

	N := int32(NTemp)

	if N < 0 {
		log.Fatal("Invalid input value.... value less than zero (0)")
	} else if N <= 100 {
		eN := N % 2
		switch {
		case N > 0 && eN > 0:
			fmt.Println("Weird")

		case (N == 2 || N <= 5) && (eN == 0):

			fmt.Println("Not Weird")

		case (N == 6 || N <= 20) && (eN == 0):
			fmt.Println("Weird")

		case N > 20 && eN == 0:
			fmt.Println("Not Weird")

		}
	}
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}
