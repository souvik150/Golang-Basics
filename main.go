package main

import (
	"fmt"
)

var l int = 23

func main() {
	fmt.Println("Hello to the world")

	fmt.Println(l)

	var i int
	i = 42

	fmt.Println(i)

	//var j int = 27
	var j float32 = 27
	fmt.Println(j)

	k := 99.
	fmt.Println(k)

	fmt.Printf("%v, %T\n", j, j)
	fmt.Printf("%v, %T\n", k, k)

	var q float32 = 42.5
	fmt.Printf("%v, %T\n", i, i)

	var w int
	w = int(q)
	fmt.Printf("%v, %T\n", w, w)

	n := 1 == 1
	m := 1 == 2

	fmt.Printf("%v, %T\n", n, n)
	fmt.Printf("%v, %T\n", m, m)

	var z bool
	fmt.Printf("%v, %T\n", z, z)

	q1 := 10
	q2 := 3

	fmt.Println(q1 + q2)
	fmt.Println(q1 - q2)
	fmt.Println(q1 * q2)
	fmt.Println(q1 / q2)
	fmt.Println(q1 % q2)

	var z1 int = 10
	var z2 int8 = 3
	fmt.Println(z1 + int(z2))

	n1 := 3.14
	n1 = 13.7e72
	n1 = 2.1E14

	fmt.Printf("%v, %T", n1, n1)

}
