package main

import "fmt"

//Golang doesnt have any OOP principle
//To create object/instance, we have to use struct.


func main(){

	type Person struct{
		name, address, phoneNumber string
		age int
	}

	myPerson := Person{name: "Nabiel", age: 21, address: "test", phoneNumber: "+62895367831730"}
	var p2 *Person = &myPerson //referencing my person
	p2.name = "dsadsa"
	fmt.Println(myPerson.name)

}