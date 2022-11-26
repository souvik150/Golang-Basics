package main

import (
	"fmt"
	"reflect"
)

type Doctor struct {
	number     int
	actorName  string
	companions []string
	episodes   []string
}

// Naming convention Uppercase is going to export
//type Doc struct {
//	Number     int
//	ActorName  string
//	Companions []string
//}

type Animal struct {
	Name   string `max:"100"`
	Origin string
}

type Bird struct {
	Animal
	SpeedKPH float32
	CanFly   bool
}

func main() {
	statePops := make(map[string]int)
	//statePops = map[string]int{
	//	"TN": 2333666,
	//	"TS": 357856,
	//	"UP": 23554653,
	//	"MP": 568045,
	//}
	statePops = map[string]int{
		"TN": 2333666,
		"TS": 357856,
		"UP": 23554653,
		"MP": 568045,
	}

	fmt.Println(statePops)
	fmt.Println(statePops["TN"])

	statePops["GJ"] = 34556234
	fmt.Println(statePops["GJ"])

	pop, ok := statePops["TO"]
	fmt.Println(pop, ok)

	_, ok = statePops["TO"]
	fmt.Println(ok)

	if q, ok := statePops["GJ"]; ok {
		fmt.Println(q)
	}

	aza, ok := statePops["TN"]
	fmt.Println(aza, ok)

	fmt.Println(len(statePops))

	sp := statePops

	//THis is going to affect not only sp but also statePops
	delete(sp, "TN")

	fmt.Println(sp)
	fmt.Println(statePops)

	aDoctor := Doctor{
		number:    3,
		actorName: "Jane Doe",
		companions: []string{
			"Souvik", "John",
		},
	}

	aDoctor1 := Doctor{
		3,
		"Jane Doe",
		[]string{
			"Souvik", "John",
		},

		//Drawback as we have to use this empty slice
		[]string{},
	}

	fmt.Println(aDoctor)
	fmt.Println(aDoctor1)

	fmt.Println(aDoctor.companions[1] + " " + aDoctor.actorName)

	b := Bird{}
	b.Name = "Emu"
	b.Origin = "Australia"
	b.SpeedKPH = 48
	b.CanFly = false

	fmt.Println(b)

	b1 := Bird{
		Animal:   Animal{Name: "Crow", Origin: "India"},
		SpeedKPH: 45,
		CanFly:   true,
	}

	fmt.Println(b1)

	t := reflect.TypeOf(Animal{})
	field, _ := t.FieldByName("Name")
	fmt.Println(field.Tag)
}
