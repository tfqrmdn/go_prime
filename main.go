package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

// IsPrime return true if int prime
func IsPrime(num int) bool {
	// num DIV 2 to add efficiency
	// https://www.thepolyglotdeveloper.com/2016/12/determine-number-prime-using-golang/
	for i := 2; i <= int(math.Floor(float64(num)/2)); i++ {
		if num%i == 0 { // mod
			return false
		}
	}
	return true
}

func main() {
	reader := bufio.NewReader(os.Stdin) // reader
	fmt.Print("input: ")
	input, _ := reader.ReadString('\n') // get input
	input = strings.TrimSpace(input)    // remove \n from string
	n, _ := strconv.Atoi(input)         // str to int

	if n < 1 {
		fmt.Println("input should >= 1")
	}

	index := 0 // index of n
	it := 2    // start of Prime number

	for index < n {
		if IsPrime(it) {
			fmt.Print(it, ",")
			it++
			index++
		} else {
			it++
		}
	}

}
