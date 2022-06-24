package main

import (
	"errors"
	"fmt"
)

func divider(n, div int) (int, error) {
	if div == 0 {
		return 0, errors.New("Divided by Zero")
	}
	return n / div, nil
}

func main() {
	a, err := divider(10, 0)
	if err != nil {
		fmt.Println("error", err)
	} else {
		fmt.Println(a)
	}
}
