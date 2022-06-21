package main

import "fmt"

type Person struct {
	key string
	val string
}

func getMyName(firstName string, lastName string) { //func with param
	fmt.Printf("%s %s", firstName, lastName)
}

func getMyNameReturningValue(firstName, lastName string) string {
	return fmt.Sprintf("\nHello %s %s!", firstName, lastName)
}

func returningMultipleValue() (string, int, Person) { //returning multiple value, new concept for me! you need to unpack the whole value tho.
	return "Nabil", 11, Person{"k", "v"}
}

func namedReturnVal(addr, bp string) (address, birthPlace string) { //named return val
	address = addr
	birthPlace = bp
	return
}

func sumAllInputtedNumber(nums ...int) (total int) { //variadic function
	total = 0

	for _, num := range nums {
		total += num
	}

	return total
}

func funcAsParam(firstParam string, firstFunc func(string) string) string {
	return fmt.Sprintf("%s", firstFunc(firstParam))
}

func sayHi(param string) string {
	return fmt.Sprintf("hello %s !", param)
}

func counter() func() { //closure
	counter := 0
	inc := func() {
		counter++
		fmt.Println(counter)
	}
	return inc
}

func main() { //Func declaration
	getMyName("Nabil", "Izzullah")

	fmt.Println(getMyNameReturningValue("Nabil", "Izzullah"))

	myName, myAge, myMap := returningMultipleValue()
	fmt.Println(myName, myAge, myMap)

	addr, bp := namedReturnVal("tes", "tes")
	fmt.Printf("%s %s\n", addr, bp)

	fmt.Println(sumAllInputtedNumber(1, 2, 3, 4, 5))

	totalSum := sumAllInputtedNumber  //function as first-class citizen
	fmt.Println(totalSum(1, 2, 5, 6)) //variable used

	fmt.Println(funcAsParam("Nabil", sayHi))

	filterWord := func(name string) bool { //anon func
		return name == "test"
	}

	fmt.Println(filterWord("test"))

	count := counter()
	count()
	count()
	count()

}
