package main

import "fmt"

//interface is an abstract data type
//usually created to define methods that will be used
//used as a contract/binding for struct

type studentInterface interface { //interface contract creation
	getStudentName() string
	getAge() int
}

type Student struct { //struct creation
	fname, lname, addr string
	age                int
}

func (s Student) getStudentName() string {
	return fmt.Sprintf("%s %s", s.fname, s.lname)
}

func (s Student) getAge() int {
	return s.age
}

func main() {
	var student1 studentInterface //we use the contract here
	student1 = Student{
		fname: "Nabil",
		lname: "Izzullah",
		addr:  "addr",
		age:   13,
	}
	fmt.Println(student1.getStudentName())
	fmt.Println(student1.getAge())
}
