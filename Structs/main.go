package main

import "fmt"

type contactInfo struct{
	email string
	zipCode int
}

type person struct {
	firstName string
	lastName  string

	// Embedded structures
	// contact  contactInfo
	contactInfo       			// this is equivalent to "contactInfo contactInfo"     
}

func main() {
	//  SYNTAX 1    (not recommended)
	// alex := person{"Alex", "Anderson"}

	//  SYNTAX 2
	// alex := person{firstName: "Alex", lastName: "Anderson"}   // second way

	//  SYNTAX 3
	jim := person{
		firstName: "Jim",
		lastName: "Party",
		contactInfo: contactInfo{
			email: "jim@gmail.com",
			zipCode: 635210,
		},
	}

	jimPointer := &jim
	jimPointer.updateName("jimmy");
	//OR
	jim.updateName("jimmy")              // Both this line and the above two lines are equivalent

	jim.print()
}


func (p person)print(){
	fmt.Printf("%+v",p)
}

func (pointerToPersion *person)updateName(newFirstName string){
	(*pointerToPersion).firstName = newFirstName
}