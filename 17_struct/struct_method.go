package main

import "fmt"

//Method is a function that can be attached into struct/other type.
//Method is function

type Student struct {
	fname, lname, addr string
	age                int
}

func (s Student) getStudentName() string { //Struct method
	return fmt.Sprintf("%s %s", s.fname, s.lname)
}

func (s *Student) changeStudentName(fname string) { //struct method with pointer, to mutate the initial struct
	s.fname = fname
}

func main() {
	stu := Student{
		fname: "Nabil",
		lname: "Izzullah",
		addr:  "addr",
		age:   21,
	}
	fmt.Println(stu.getStudentName())
	stu.changeStudentName("ddd")
	fmt.Println(stu.fname)
}
