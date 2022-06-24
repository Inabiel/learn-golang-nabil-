package main

import "fmt"

//if we want to mutate a variable from outside (from func,) we need to use pointer func
//because golang is pass-by-value language, if u dont use the reference you wont change the initial var
//this is because when u run func without ref/deref, it will copy the value into new var

func incrementAge(i *int) {
	*i += 1
}

func main() {
	age := 12
	incrementAge(&age)
	fmt.Println(age)
}
