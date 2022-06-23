package main

import "fmt"

// Defer is used to let the compiler know that the deferred func is ran after every other script is done.
// it is placed on the last stack. if you run more than 1 deferred func it is ran last in first out! be careful.
// this happens because the deferred statement is stacked on top of the previous one, and called from the top.
//pretty unique, huh?

func defer1() {
	fmt.Println("1")
}

func defer2() {
	fmt.Println("2")
}

func main() {
	defer func() {
		fmt.Println("3") //this runs last
	}() //example of defer on IIFE
	fmt.Println("Mulai")
	defer defer2() //this runs second
	fmt.Println("Selesai")
	defer defer1() //this is first!
}
