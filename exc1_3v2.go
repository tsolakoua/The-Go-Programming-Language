// Takes as input a number from the cmd which is string)
// converts it to integer and makes a slice with as many elements as the length of the input number
// each index of the slice contains the index number as value
// converts the values of the slice to string and echoes the values in two different methods, comparing running time

package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {

	args_input := os.Args[1]
	max := stringToInt(args_input)
	numSlice := createSlice(max)

	echo1(numSlice)
	echo2(numSlice)
}

func createSlice(max int) []string {

	numSlice := make([]string, max)
	for i := 1; i < max; i++ {
		j := IntToString(i)
		numSlice[i] = j
	}

	return numSlice
}

func stringToInt(args_input string) int {

	i, err := strconv.Atoi(args_input)
	fmt.Println(err)

	return i
}

func IntToString(args_input int) string {

	i := strconv.Itoa(args_input)

	return i

}

func echo1(numSlice []string) {

	var result, sep string
	start := time.Now()

	for i := 0; i < len(numSlice); i++ {

		result += sep + numSlice[i]
		sep = " "
	}
	elapsed := time.Since(start)
	fmt.Println(result)
	fmt.Println("Time took", elapsed)

}

func echo2(numSlice []string) {
	start := time.Now()
	result := strings.Join(numSlice[0:], " ")
	elapsed := time.Since(start)
	fmt.Println(result)

	fmt.Println("Time took", elapsed)
}
