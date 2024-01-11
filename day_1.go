package main

import (
	"fmt"
	"math"
)

func main() {
	var name1 = "Sadanand"      //type is infreed
	var name2 string = "gadwal" //type is string

	var add string = name1 + " " + name2
	x := 10
	y := 20
	z := x + y
	q := x - y
	p := x * y
	r := x / y
	s := math.Pow(2, 3)

	//multiple variable
	var a, b, c, d int = 1, 3, 5, 7

	// Go Variable Naming Rules
	var myVariableName = "myVariableName"
	var MyVariableName = "MyVariableName"
	var my_variable_name = "my_variable_name"

	// output
	//string
	fmt.Println(name1)
	fmt.Println(name2)
	fmt.Println(add)
	//int
	fmt.Println("\n", "Integers Print")
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println("\n", "Integers Math Operations Print + , - , / , * ,square ")
	fmt.Println("Addition", z)
	fmt.Println("Multiplication", p)
	fmt.Println("Substract", q)
	fmt.Println("Divide", r)
	fmt.Printf("Square %.1f", s)
	//multiple
	fmt.Println("\n", "multiple variable output Print")
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println("\n", "Naming Rules :")
	fmt.Println(myVariableName)
	fmt.Println(MyVariableName)
	fmt.Println(my_variable_name)

}
