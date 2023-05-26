// format printing using printf
package main

import "fmt"

func main() {
	// const name = "Priyanka"
	// const age = 28
	// or
	const name, age = "Priyanka", 28
	// fmt.Printf("%s is %d years old.\n", name, age)
	fmt.Printf("%v is %v years old, and first value is of %T type and second value is of %T type", name, age, name, age)
}
