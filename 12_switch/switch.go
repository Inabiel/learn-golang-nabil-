package main

import "fmt"

func main() {
	name := "Bambang"

	switch name {
	case "Bambang":
		fmt.Println("Nama bambang")
	default:
		fmt.Println("Nama bukan bambang")
	}

	//short statement switch

	switch length := len(name); length > 2{
	case true:
		fmt.Println("Benar")
	case false:
		fmt.Println("Salah")
	}

	//non-conditional switch
	length := len(name)
	switch{
	case length == 1:
		fmt.Println("Panjang 1")
	default:
		fmt.Println("Panjang lebih dari 1")
	}
}