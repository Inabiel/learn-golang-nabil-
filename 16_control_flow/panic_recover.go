package main

import "fmt"

//Panic and recover is a method on golang to throw an exceptions.



func main(){
	defer func(){
		fmt.Println("deferred func here")
		if r:= recover(); r!= nil{ //recover only works on deferred func, because it is run last, then it can recover the panic.
			fmt.Println("it doesnt panic")
		}
	}()

	panic("Boom") // Use it to throw an error, similar to throw() on other language
}

//When to use defer, recover, panic?
//Use defer to run any process last,
//use panic to throw any exception,
//use recover to catch exception and run the process forward.