package main

import "fmt"

//we dont have any in golang
//to use any type, we can use interface{}

type any interface{} // we create type called any that is interface{} underneath

func main() {
	var seluruhType any
	seluruhType = 1
	seluruhType = true
	seluruhType = "a" //its not causing error when we change the type!
	//keep in mind this is a bad practice, because we can't really ensure the type provided.
	//why use static-typed language when all u need is any, right? just use py or something.
	fmt.Printf("%s", seluruhType)
}
