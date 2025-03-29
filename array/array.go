package array

import (
	"fmt"
)

type Array struct {
	data     []int
	size     int
	capacity int
}

type ArrayInterface interface{
	ReturnValue () int 
	Error () error
}

func NewArray(capacity int) *Array {
	return &Array{
			data:     make([]int, capacity),
			size:     0,
			capacity: capacity,
	}
}

func (a *Array) FindMax()(int,int,error){
	if  a.size == 0{
		return 0,0,fmt.Errorf("array is empty")
	}
	max := a.data[0]
	index := 0
	for index:= 1;index<a.size;index++{
		if a.data[index] > max{
			max = a.data[index]
		}
	}
	return max,index,nil
}

func (a *Array) FindMin() (int,int,error) {
	if a.size == 0 {
		return 0,0,fmt.Errorf("array is empty")
	}
	min := a.data[0]
	index := 0
	for index = 1; index < a.size; index++ {
		if a.data[index] < min {
			min = a.data[index]
		}
	}
	return min,index,nil
}


func (a *Array) FindMaxMin() (int,int,[]int,error) {
	if a.size == 0 {
		return 0,0,[]int{}, fmt.Errorf("array is empty")
	}
	max := a.data[0]
	min := a.data[0]
	maxIndex , minIndex := 0,0
	for i := 1; i < a.size; i++ {
		if a.data[i] > max {
			max = a.data[i]
			maxIndex = i
		}
		if a.data[i] < min {
			min = a.data[i]
			minIndex = i
		}
	}
	return max,min,[]int{maxIndex,minIndex}, nil
}


