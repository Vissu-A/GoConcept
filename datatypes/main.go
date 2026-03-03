package main

// standard library for formatted I/O, similar to java's System.out.println and "stdio.h" in C.
import (
	"demo/datatypes/aggdatatypes"
	"fmt"
)

// main method like "static void main()" in java and "__main__" in python
func main() {
	println("Hello World..!")
	fmt.Println("Welocome to Go..!")
	aggdatatypes.ArrayExample()
	aggdatatypes.SliceExample()
}
