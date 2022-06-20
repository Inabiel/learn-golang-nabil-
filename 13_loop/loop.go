package main

import "fmt"

func main() {

	counter := 1

	for counter <= 10 {
		fmt.Println(counter)
		counter++
	}

	//js-like for loops

	for i := 11; i <= 20; i++{
		fmt.Println(i)
	}

	//range(array/slice) loops
	arr := [...]int{1,4,5,6}

	for k,v := range arr{
		fmt.Printf("{%d,%d}\n",k,v)
	}

	//iterating over map
	mapOfLanguage := map[string]string{
		"Indonesia":"IN",
		"English":"EN",
		"France":"FR",
		"Italia":"IT",
	}
	for k,v := range mapOfLanguage{
		fmt.Printf("{'Language:%s, Abbreviated:%s'}, ", k, v)
	}

}