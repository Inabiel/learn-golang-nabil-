package main

import "fmt"

//golang is passed-by-value at default
//it means if you copy an object into other object, it only copy the value not the ref

type Address struct {
	Province, City, Street string
}

func main() {

	firstAddress := Address{
		Province: "Jogja",
		City:     "Bantul",
		Street:   "Imogiri",
	}

	secondAddress := firstAddress            // we copy the object
	secondAddress.City = "Gunkid"            // we change the city
	fmt.Println(firstAddress, secondAddress) //it does differs

	thirdAddress := &firstAddress
	*thirdAddress = Address{"Jogja", "Sleman", "Imogiri"} // we mutate the struct of firstAddress
	fmt.Println(firstAddress, secondAddress, *thirdAddress)

	fourthAddress := new(Address) //new keyword, creating empty pointer that addressed into a struct
	*fourthAddress = Address{
		Province: "DKI Jakarta",
		City:     "Jaksel",
		Street:   "Sudirman",
	}
	fmt.Println(*fourthAddress)
}
