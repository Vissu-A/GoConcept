package loop

import (
	"fmt"
)

func ForLoopExample() {
	var colourCodes map[string]string

	colourCodes = map[string]string{
		"red":   "#FF0000",
		"green": "#00FF00",
		"blue":  "#0000FF",
	}

	for colour, code := range colourCodes {
		fmt.Println(colour, ":", code)
	}

}
