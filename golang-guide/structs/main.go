package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

type contactInfo struct {
	email   string
	zipCode int
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

	fmt.Printf("%+v\n", notAlex)

	jim := person{
		firstName: "Jim",
		lastName: "Harper",
		contact: contactInfo{
			email: "jim.harper@dunder.mifflin.com",
			zipCode: 94000,
		},
	}

	fmt.Printf("%+v", jim)
}
