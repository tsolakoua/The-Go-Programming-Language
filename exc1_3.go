package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {

	echo1()
	echo2()
}

func echo1() {

	var s, sep string
	start := time.Now()

	for i := 1; i < len(os.Args); i++ {

		s += sep + os.Args[i]
		sep = " "
	}
	elapsed := time.Since(start)
	fmt.Println(s)
	fmt.Println("Time took", elapsed)

}

func echo2() {
	start := time.Now()
	result := strings.Join(os.Args[1:], " ")
	elapsed := time.Since(start)
	fmt.Println(result)

	fmt.Println("Time took", elapsed)
}
