package main

import (
	"github.com/Orishigami/go-network2/work1"

	"fmt"
)

type Person struct {
	name   string
	age    int
	job    string
	salary int
}

func main() {
	message := work1.Work1("Gladys")
	fmt.Println(message)

	var student1 string = "John" //type is string
	var student2 = "Jane"        //type is inferred
	x := 2                       //type is inferred

	fmt.Println(student1)
	fmt.Println(student2)
	fmt.Println(x)

	var a string
	var b int
	var c bool

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

	var d, e, f, g int = 1, 3, 5, 7

	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)
	fmt.Println(g)

	var h bool = true    // Boolean
	var i int = 5        // Integer
	var j float32 = 3.14 // Floating point number
	var k string = "Hi!" // String

	fmt.Println("Boolean: ", h)
	fmt.Println("Integer: ", i)
	fmt.Println("Float:   ", j)
	fmt.Println("String:  ", k)

	var arr1 = [3]int{1, 2, 3}
	arr2 := [5]int{4, 5, 6, 7, 8}

	fmt.Println(arr1)
	fmt.Println(arr2)

	arr3 := [6]int{10, 11, 12, 13, 14, 15}
	myslice := arr3[2:4]

	fmt.Printf("myslice = %v\n", myslice)
	fmt.Printf("length = %d\n", len(myslice))
	fmt.Printf("capacity = %d\n", cap(myslice))

	myslice1 := []int{}
	fmt.Println(len(myslice1))
	fmt.Println(cap(myslice1))
	fmt.Println(myslice1)

	myslice2 := []string{"Go", "Slices", "Are", "Powerful"}
	fmt.Println(len(myslice2))
	fmt.Println(cap(myslice2))
	fmt.Println(myslice2[2])
	fmt.Println(myslice2)

	var l = 15 + 25
	fmt.Println(l)

	var (
		sum1 = 100 + 50    // 150 (100 + 50)
		sum2 = sum1 + 250  // 400 (150 + 250)
		sum3 = sum2 + sum2 // 800 (400 + 400)
	)
	fmt.Println(sum3)

	var m = 10
	m += 5
	fmt.Println(m)

	var pers1 Person

	pers1.name = "Hege"
	pers1.age = 45
	pers1.job = "Teacher"
	pers1.salary = 6000

	fmt.Println("Name: ", pers1.name)
	fmt.Println("Age: ", pers1.age)
	fmt.Println("Job: ", pers1.job)
	fmt.Println("Salary: ", pers1.salary)

}
