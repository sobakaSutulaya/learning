package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	contactInfo
}

type contactInfo struct {
	email   string
	zipCode int
}

func (p person) print() {
	fmt.Printf("%+v", p)
}

// it is will not work cause Go creates a copy of real value and updated it
// so : 
// p := person{firstName: "name", lastName: "lName"}
// p.changeName("new name")
// and after that p will have an old name cause of pass by value shit
// so we should use pointers like in method below
// func (p person) changeName(newName string) {
// 	p.firstName = newName
// }

func (pointerToPerson *person) changeName(newName string) {
	//pointer to pointer is an actual value
	(*pointerToPerson).firstName = newName
}

func main() {
	// bad way, if a want to swap fields in struct for some reason, this just will work fine but i'll have firstname=Andersen and lastname=Alex
	// alex := person{"Alex", "Andersen"}
	// it is good? kind of yes
	alex := person{firstName: "Alex", lastName: "Andersen"}
	fmt.Println(alex)

	//can create empty and set fields mannually
	var notAlex person
	fmt.Printf("%+v\n", notAlex)

	notAlex.firstName = "notAlex"
	notAlex.lastName = "definetly not Andersen"

	// fmt.Printf("%+v\n", notAlex)

	jim := person{
		firstName: "Jim",
		lastName:  "Harper",
		contactInfo: contactInfo{
			email:   "jim.harper@dunder.mifflin.com",
			zipCode: 94000,
		},
	}
	
	//jimPointer is a memory address of jim
	jimPointer := &jim
	jimPointer.changeName("Jimmy")
	jim.print()
}
