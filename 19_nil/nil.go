package main

import "fmt"

//nil is a default provided by golang, it is same as NULL on some other programming lang
//nil is only applicable on func, map, interface{}, slice, pointer(*), and channel

func createNewMap(n string) map[string]interface{} {
	if n == "" {
		return nil
	}
	return map[string]interface{}{
		"name": n,
	}
}

func main() {
	a := createNewMap("")     // check using no value
	fmt.Println(a)            // it returns nil/map[]
	b := &a                   //referencing a
	*b = createNewMap("test") //setting pointer into a, running func with val
	fmt.Println(*b)           //see the result
}
