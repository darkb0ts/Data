package main

import (
	"fmt"

	"github.com/darkb0ts/Data/array" // Correct import path
)

func main() {
	arr := []int{1, 2, 3, 4, 5}
	value,value1, err := array.Max(arr)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	value7,err := array.Random(20,-20,10,true)
	if err != nil {
		fmt.Println("Error:", err)
		return
	} 
	fmt.Println("Random:", value7, err)
	fmt.Println("Array:", value, value1, err)
}

