package main

import (
	"fmt"
	//"math"
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
	n1 = 2.1e14

	fmt.Printf("%v, %T\n", n1, n1)

	var n2 complex64 = 1 + 2i
	n2 = 1i
	fmt.Printf("%v, %T\n", n2, n2)

	var n3 complex64 = 2 + 4i
	var n4 complex64 = 5 + 6i

	fmt.Println(n3 + n4)
	fmt.Println(n3 - n4)
	fmt.Println(n3 * n4)
	fmt.Println(n3 / n4)

	var n5 complex128 = 1 + 2i
	fmt.Printf("%v, %T\n", real(n5), real(n5))
	fmt.Printf("%v, %T\n", imag(n5), imag(n5))

	s := "this is a string"
	fmt.Printf("%v, %T\n", s[2], s[2])
	fmt.Printf("%v, %T\n", string(s[2]), s[2])

	s1 := "this is also a string"
	fmt.Printf("%v, %T\n", s+" "+s1, s+s1)

	s2 := "this is a string"
	b1 := []byte(s2)

	fmt.Printf("%v, %T\n", b1, b1)

	//rune
	r := 'a'
	fmt.Printf("%v, %T\n", r, r)

	var r1 rune = 'b'
	fmt.Printf("%v, %T\n", r1, r1)

	//	The below is not allowed
	//	const myConst float64 = math.Sin(1.57)

	const myConst float64 = 32
	fmt.Printf("%v, %T\n", myConst, myConst)

	const m1 = 34
	var b2 int16 = 27
	//b2 := 65
	fmt.Printf("%v, %T\n", m1+b2, m1+b2)

}
