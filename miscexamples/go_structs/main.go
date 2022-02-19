package main

import "fmt"

type contactInfo struct {
	email string
	zip   int
}

type person struct {
	firstName string
	lastName  string
	//contact   contactInfo
	contactInfo
}

func main() {
	fmt.Println("Struct demos")
	//Order of fields assignment is this crazy!!
	// alex := person{"Alix", "Anderson"}
	alex := person{firstName: "Raja", lastName: "Ram"}

	fmt.Println(alex)

	var tarzanPerson person

	fmt.Printf("%+v", tarzanPerson)

	tarzanPerson.firstName = "Joe"
	fmt.Println(tarzanPerson)
	fmt.Printf("%+v", tarzanPerson)

	jim := person{
		firstName: "Jim",
		lastName:  "Stone",
		// contact: contactInfo{
		// 	email: "jim@google",
		// 	zip:   50000,
		// }
		contactInfo: contactInfo{
			email: "jim@google",
			zip:   50000,
		},
	}

	//fmt.Printf("%+v", jim)
	//Everything is pass by value in Golang so this would not udpate jim
	jim.updateName("Bikram")
	jim.print()

	fmt.Println("Pointer update =========")
	//jimPointer := &jim
	//jimPointer.updateNameByPointer("Kajol")
	//Or we can also call updateNameBypointer since receiver has pointer to function
	jim.updateNameByPointer("AAJA")
	jim.print()
}
func (p *person) updateNameByPointer(firstName string) {
	(*p).firstName = firstName

}

func (p person) updateName(firstName string) {
	p.firstName = firstName

}

func (p person) print() {
	fmt.Println("print the receiver person")
	fmt.Printf("%+v", p)
}
