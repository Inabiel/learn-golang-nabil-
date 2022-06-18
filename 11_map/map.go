package main

import "fmt"

func main() {
	obj := map[string]string{
		"fname": "Nabiel",
		"lname": "Izzullah",
	} //creating map with string-string k-v

	fmt.Println(obj["fname"])

	obj2 := map[string]interface{}{
		"fname":"Nabiel",
		"lname":"Izzullah",
		"age":21,
	} //creating map with string-dynamic k-v
		
	obj2["telp"] = "0895367831730"
	fmt.Println(obj2["telp"])

}