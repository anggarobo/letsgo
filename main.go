package main

import "fmt"

func main() {
	// Non Decimal
	var positveNum uint8 = 99
	var negativeNum = -1243423644

	// Decimal
	var phi = 3.14

	// Boolean
	var isExists bool = true

	// String
	var greeting string = "good evening"
	var aboutMe = `
	Name : John Doe
	Job  : Coder
	Role : Backend Engineer
	`

	fmt.Println("Lets Go!")
	fmt.Printf("Positive Number => %d\n", positveNum)
	fmt.Printf("Negative Number => %d\n", negativeNum)
	fmt.Printf("Decimal Number => %f\n", phi)
	fmt.Printf("Phi => %.2f\n", phi)
	fmt.Printf("isExists = %t \n", isExists)
	fmt.Printf("Hi, %s everyone! \n", greeting)
	fmt.Println(aboutMe)
}
