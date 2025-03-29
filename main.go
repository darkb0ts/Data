package main

import (
	"fmt"

	"github.com/darkb0ts/Data/array" // Correct import path
)

func main() {
	arr := array.NewArray(10)
	fmt.Println("Array created with capacity:", arr)
}