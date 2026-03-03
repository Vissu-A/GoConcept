package aggdatatypes

import "fmt"

func StructExample() {

	// Define struct
	type user struct {
		username string
		userid   int
		gender   string
		balance  float64
	}

	tony := user{"tony", 1, "male", 10000.00}
	captian := user{username: "captian", userid: 2, gender: "male", balance: 15000.00}

	fmt.Println(tony)
	fmt.Println(captian)

}
