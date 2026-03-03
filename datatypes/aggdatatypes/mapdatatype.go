package aggdatatypes

import (
	"fmt"
)

func MapExample() {
	var stuMarks map[string]int

	stuMarks = map[string]int{
		"John": 85,
		"Jane": 92,
		"Bob":  78,
	}
	fmt.Println(stuMarks)

	// creating map using make function
	teachers := make(map[string]string)
	fmt.Println("Empty Teachers map: ", teachers)

	// adding value to the map
	teachers["Math"] = "Mr. Smith"
	teachers["Science"] = "Ms. Johnson"
	teachers["Physics"] = "tony"
	teachers["History"] = "Ms. Davis"
	fmt.Println("Teachers map after adding values: ", teachers)

	// Accessing value from the map
	fmt.Println("Teacher for Math: ", teachers["Math"])
	fmt.Println("Physics teacher: ", teachers["Physics"])

	// Deleting a key-value pair from the map using built-in delete function
	delete(teachers, "History")

	fmt.Println("Teachers map: ", teachers)

}
