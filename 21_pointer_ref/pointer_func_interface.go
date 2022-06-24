package main

import "fmt"

type KTP struct {
	name, religion, phoneNum string
	age                      int
}

type KTPInterface interface {
	getName() string
	getReligion() string
	getPhoneNum() string
	getAge() int
	mutateName(s string)
}

// getAge implements KTPInterface
func (k *KTP) getAge() int {
	return k.age
}

// getName implements KTPInterface
func (k *KTP) getName() string {
	return fmt.Sprintf("%s", k.name)
}

// getPhoneNum implements KTPInterface
func (k *KTP) getPhoneNum() string {
	return fmt.Sprintf("%s", k.phoneNum)
}

// getReligion implements KTPInterface
func (k *KTP) getReligion() string {
	return fmt.Sprintf("%s", k.religion)
}

// mutateName implements KTPInterface
func (k *KTP) mutateName(s string) {
	k.name = s
}

func main() {
	var myKTP KTPInterface = &KTP{
		name:     "Nabil Izzullah",
		religion: "Moslem",
		phoneNum: "+6289xxxxxx",
		age:      21,
	}
	myKTP.mutateName("Saleh")
	fmt.Println(myKTP.getReligion())
	fmt.Println(myKTP.getName())
	fmt.Println(myKTP)
}
