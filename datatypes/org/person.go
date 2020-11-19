package org

type identifiable interface {
	ID() string
}

type Person struct {
	fisrtName string
	lastName  string
}

func NewPerson(fisrtName, lastName string) Person {
	return Person{
		fisrtName: fisrtName,
		lastName:  lastName,
	}

}
func (p Person) ID() string {
	return "12345"
}
