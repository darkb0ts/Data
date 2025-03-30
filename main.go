package main

import (
	"fmt"

	"github.com/darkb0ts/Data/array" // Correct import path
)

func main() {
	arr,_ := array.Random(20,1,100)
	fmt.Println("Array:", arr)
	value,_:=array.Sum(arr)
	fmt.Println("Reversed Array:", value)
}

