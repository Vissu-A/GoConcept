package aggdatatypes

import (
	"fmt"
	"reflect"
)

func ArrayExample() {

	var stuRoolNo [4]int                   // Creating/Defining/Instantiating the array of fixed size 4.
	stuRoolNo = [4]int{001, 002, 003, 004} // Why we need to specify the size and type of array again while intializing.

	stuGrades := [4]string{"A", "C", "D", "B"}

	fmt.Println("Student Rool No: ", stuRoolNo)
	fmt.Println("Student grades: ", stuGrades)
	fmt.Println("Student Rool No: ", reflect.TypeOf(stuRoolNo))
	fmt.Println("Student grades: ", reflect.TypeOf(stuGrades))
}
