package main

import (
	"fmt"
)

func main() {

	// copy by value, a and b are different variables with same value.
	fmt.Println("================Copy by value====================")
	a := 27
	b := a

	fmt.Println("Variable A value: ", a)
	fmt.Println("Variable B value: ", b)

	a = 52
	fmt.Println("Variable A value: ", a)
	fmt.Println("Variable B value: ", b)

	// copy by address.
	fmt.Println("================Copy by address===================")
	msg := "hello"
	greeting := &msg

	fmt.Println("Variable msg value: ", msg)
	fmt.Println("Variable msg memory location: ", &msg)
	fmt.Println("Variable greeting value: ", greeting) // This will print the memory location address.
	fmt.Println("We can see both variables 'msg' and 'greeting' are having same memory location")

	msg = "welcome" // Overriding variable "msg" value.

	fmt.Println("Variable msg value: ", msg)

	// De-referencing pointer to access the actual value it's points to.
	fmt.Println("Variable greeting value: ", *greeting)
}
