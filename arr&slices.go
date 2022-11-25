package main

import "fmt"

func main() {
	//grades := [3]int{97, 85, 93}
	grades := [...]int{97, 85, 93}
	fmt.Printf("Grades: %v, %v, %v\n", grades[0], grades[1], grades[2])
	fmt.Printf("Grades: %v\n", grades)

	var students [3]string
	fmt.Printf("Students: %v\n", students)
	students[0] = "Lisa"
	students[1] = "Ahmed"
	students[2] = "Arnold"

	fmt.Printf("Student #1: %v\n", students[1])
	fmt.Printf("Length of the Student Array is: %v\n", len(students))

	var identityMatrix [3][3]int = [3][3]int{[3]int{1, 0, 0}, [3]int{0, 1, 0}, [3]int{0, 0, 1}}
	fmt.Println(identityMatrix)

	fmt.Println("======= Pointers =========")
	//	Pointers
	a := [...]int{1, 2, 3}
	b := &a
	b[1] = 5

	c := a
	//Changes only c
	c[2] = 4

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

	fmt.Println("======= Slices =========")
	//	Slice
	q := []int{1, 2, 3}
	w := q

	//Changes both q and w
	w[1] = 5

	fmt.Println(q)
	fmt.Println(w)
	fmt.Printf("Length: %v\n", len(q))
	fmt.Printf("Capacity: %v\n", cap(q))

	aq := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	p := aq[:]
	o := aq[3:]
	i := aq[:6]
	u := aq[3:6]

	aq[5] = 42

	fmt.Println(aq)
	fmt.Println(p)
	fmt.Println(o)
	fmt.Println(i)
	fmt.Println(u)

	aw := make([]int, 3, 100)
	fmt.Println(aw)
	fmt.Printf("Length: %v\n", len(aw))
	fmt.Printf("Capacity: %v\n", cap(aw))

	a1 := []int{}
	fmt.Println(a1)
	fmt.Printf("Length: %v\n", len(a1))
	fmt.Printf("Capacity: %v\n", cap(a1))

	a1 = append(a1, 1, 2)
	fmt.Println(a1)
	fmt.Printf("Length: %v\n", len(a1))
	fmt.Printf("Capacity: %v\n", cap(a1))

	//Adding elements to a slice
	//a1 = append(a1, 2, 3, 4, 5, 6)

	//Concatenating one slice with another
	a1 = append(a1, []int{2, 3, 4, 5, 6}...)
	fmt.Println(a1)
	fmt.Printf("Length: %v\n", len(a1))
	fmt.Printf("Capacity: %v\n", cap(a1))

	//Remove the first element
	a1 = []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println(a1)

	b1 := a1[1:]
	fmt.Println(b1)

	//Removing the last element
	b2 := a1[:len(a1)-1]
	fmt.Println(b2)

	//Removing the second element
	b3 := append(a1[:2], a1[3:]...)
	fmt.Println(b3)
	fmt.Println(a1)

}
