package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello world")
	var name string = "Manthan Bhatt"
	age := 25
	var income bool
	fmt.Println(name)
	fmt.Println(age)
	fmt.Println(income)
	if age == 25 {
		fmt.Println("prime")
	}

	const PI float32 = 3.14
	fmt.Println(PI)
	multiVariable()
	declareArray()
}

func multiVariable() {
	var a, b, c, d int = 1, 3, 5, 7

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)

	// confusing code
	// var e, f = 6, "Hello"
	// g, h := 7, "World!"

	// fmt.Println(e)
	// fmt.Println(f)
	// fmt.Println(g)
	// fmt.Println(h)

	var (
		j int
		k int    = 1
		l string = "hello"
	)

	fmt.Println(j)
	fmt.Println(k)
	fmt.Println(l)
}

func declareArray() {
	var arr = [...]int{1, 2, 3, 4, 5, 6} // declaring array with inferred length
	fmt.Println(arr)
	fmt.Println(arr[2])
	fmt.Println(len(arr))
	fmt.Println(arr[1:3]) // [2,3]
	fmt.Println(append(arr[1:3], 10, 20))

	myslice1 := []int{1, 2, 3}
	myslice2 := []int{4, 5, 6}
	myslice3 := append(myslice1, myslice2...)
	fmt.Println(myslice3)
}

/**
---- Variable
- var can be used anywhere
- := can only be used inside of a function
- A variable name must start with a letter or an underscore character (_)
- A variable name cannot start with a digit
- A variable name can only contain alpha-numeric characters and underscores (a-z, A-Z, 0-9, and _ )
- Variable names are case-sensitive (age, Age and AGE are three different variables)
- There is no limit on the length of the variable name
- A variable name cannot contain spaces
- The variable name cannot be any Go keywords

- No ternary operators in go lang

*/
