package main

import "fmt"

func main() {
	arrayOfPeople := map[string]string{
		"IN": "Insigna Nodu",
		"FA": "Fabrizio Armani",
	}

	for k, v := range arrayOfPeople {
		if k == "FA" {
			fmt.Println("Found FA")
			break
		}
		fmt.Println(k,v)
	}
}