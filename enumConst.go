package main

import (
	"fmt"
	//"math"
)

const (
	c7 = iota
	c8 = iota
	c9 = iota
)

const (
	c6 = iota
)

const (
	_ = iota + 4
	catSpecialist
	dogSpecialist
	snakeSpecialist
)

const (
	isAdmin = 1 << iota
	isHeadquarters
	canSeeFinancials
	canSeeAfrica
	canSeeAsia
	canSeeEurope
	canSeeNorthAmerica
	canSeeSouthAmerica
)

func main() {
	//	Enumerating consts
	fmt.Printf("%v\n", c7)
	fmt.Printf("%v\n", c8)
	fmt.Printf("%v\n", c9)

	fmt.Printf("%v\n", c6)

	var specialistType int
	fmt.Printf("%v\n", specialistType == catSpecialist)
	fmt.Printf("%v\n", catSpecialist)
	fmt.Printf("%v\n", dogSpecialist)
	fmt.Printf("%v\n", snakeSpecialist)

	var roles byte = isAdmin | canSeeFinancials | canSeeEurope
	fmt.Printf("%b\n", roles)
	fmt.Printf("Is Admin? %v\n", isAdmin&roles == isAdmin)
	fmt.Printf("Can see asia? %v", canSeeAsia&roles == canSeeAsia)

}
