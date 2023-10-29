package main

import "fmt"

func main() {
	// Inline Comment
	/*
		Multiline Comment
	*/
	// VARIABLE DECLARATION
	var firstName string = "Angga" // => with datatype
	lastName := "" // => without datatype
	lastName = "Prabowo" // => reassign

	// MULTIVARIABLE DECLARATION
	var first, second, third string
	first, second, third = "first", "second", "third"
	var fourth, fifth, sixth string = "fourth", "fifth", "sixth"
	seventh, eighth, ninth := "seventh", "eighth", "ninth"
	var numericType int8 = 1
	var booleanType bool = true
	var decimalType float32 = 5.5
	var stringType string = "string"
	_ = "underscore allow a variable is not used"
	greeting, _ := "good morning!", "underscore var is predefined but not for this"

	fmt.Println("Hello folk,")
	fmt.Printf("Wellcome back %s %s! \n", firstName, lastName)
	fmt.Printf("%s %s %s %s %s %s %s %s %s \n", first, second, third, fourth, fifth, sixth, seventh, eighth, ninth)
	fmt.Println("numericType =>", numericType)
	fmt.Println("booleanType =>", booleanType)
	fmt.Println("decimalType =>", decimalType)
	fmt.Println("stringType =>", stringType)
	fmt.Println(greeting)

}
