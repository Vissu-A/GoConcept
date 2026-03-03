package main

import (
	"fmt"
)

func main() {
	var clsTeacherName string // Default/zero value of string is "" (empty string)
	var clsGrade int          // Default/zero value of int is "0"
	var status bool           // Default/zero value of boolean is "false"
	var name string = "Tony"  // type specified explictly.
	var score = 92            // type inferred by the compiler.
	grade := "A"              // short variable declaration, only for local variables, type inferred by the compiler.

	fmt.Println(clsTeacherName)
	fmt.Println(clsGrade)
	fmt.Println(status)
	fmt.Println(name, score, grade)
}
