package main

import (
	"datatypes/org"
)

func main() {
	p := org.NewPerson{firstName: "james", lastName: "parelta"}
	println(p.ID())
}
