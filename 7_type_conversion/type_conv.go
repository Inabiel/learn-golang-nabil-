package main

import (
	"fmt"
	"strconv"
)

func main() {
	var number int = 123

	var stringedNumber string = strconv.Itoa(number)

	fmt.Println(stringedNumber)
}