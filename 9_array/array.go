package main

import "fmt"

func main() {
	var listNama = [4]string{"a", "b", "c"}
	listNama[3] = "2"

	if(len(listNama) > 3) {
		fmt.Println("Panjang!")
	}

	fmt.Printf(fmt.Sprintf("Length is %d with the data is %v", len(listNama), listNama))
}