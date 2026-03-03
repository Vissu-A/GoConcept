package aggdatatypes

import (
	"fmt"
	"slices"
)

func SliceExample() {

	// Decalraing/Defining Slice varaible "s" with no fixed size of type integer.

	var marks []int
	// This create a slice with name "s", but the value it holds is "nil". nil indicates that the variable is declared but not initalized. nil is the default values for pointers and reference types like slice if not initalized with values.
	marks = []int{95, 75, 34, 48, 92, 84}

	fmt.Println("Marks: ", marks)
	fmt.Println("Mark at index 2: ", marks[2])

	marks[4] = 97
	fmt.Println("Updated Marks: ", marks)

	marks = append(marks, 88, 74, 91)
	fmt.Println("Marks after appending: ", marks)

	// Deleting an element from the slice (removing the elements at index 3, 4)
	marks = slices.Delete(marks, 3, 5)
	fmt.Println("Marks after deletion: ", marks)

	// Copying a slice - copy by reference.
	copiedMarksRef := marks
	fmt.Println("Copied Marks by Reference: ", copiedMarksRef)

	copiedMarksVal := slices.Clone(marks)
	fmt.Println("Copied Marks by Value: ", copiedMarksVal)
}
