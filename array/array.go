package array

import (
	"fmt"
	"math/rand"
	"time"
)

type Array struct {
	data     []int
	size     int
	capacity int
}


func NewArray(capacity int) *Array {
	return &Array{
			data:     make([]int, capacity),
			size:     0,
			capacity: capacity,
	}
}

func (a *Array) Get(index int) (int, error) {
	if index < 0 || index >= a.size {
		return 0, fmt.Errorf("index out of range: %d", index)
	}
	return a.data[index], nil
}

func (a *Array) Set(index int, value int) error {
	if index < 0 || index >= a.size {
		return fmt.Errorf("index out of range: %d", index)
	}
	a.data[index] = value
	return nil
}

func (a *Array) Size() int {
	return a.size
}
func (a *Array) Capacity() int {
	return a.capacity
}
func (a *Array) IsEmpty() bool {
	return a.size == 0
}
func (a *Array) IsFull() bool {
	return a.size == a.capacity
}
func (a *Array) Append(value int) (bool, error) {
	if a.size == a.capacity {
		return false, fmt.Errorf("array is full")
	}
	a.data[a.size] = value
	a.size++
	return true, nil
}

func (a *Array) Remove(index int) (int, error) {
	if index < 0 || index >= a.size {
		return 0, fmt.Errorf("index out of range: %d", index)
	}
	if a.size == 0 {
		return 0, fmt.Errorf("array is empty")
	}

	value := a.data[index]
	for i := index; i < a.size-1; i++ {
		a.data[i] = a.data[i+1]
	}
	a.data[a.size-1] = 0 
	a.size--
	return value, nil
}

func (a *Array) Clear() {
	for i := 0; i < a.size; i++ {
		a.data[i] = 0 
	}
	a.size = 0
}

func (a *Array) Print() {
	for i := 0; i < a.size; i++ {
		fmt.Printf("%d ", a.data[i])
	}
	fmt.Println()
}

func (a *Array) Reverse() {
	for i, j := 0, a.size-1; i < j; i, j = i+1, j-1 {
		a.data[i], a.data[j] = a.data[j], a.data[i]
	}
}

func (a *Array) BubbleSort() ([]int, error) {
	if a.size == 0 {
		return nil, fmt.Errorf("array is empty")
	}
	for i := 0; i < a.size-1; i++ {
		for j := 0; j < a.size-i-1; j++ {
			if a.data[j] > a.data[j+1] {
				a.data[j], a.data[j+1] = a.data[j+1], a.data[j]
			}
		}
	}
	return a.data, nil
}

func (a *Array) Sum() (int,error){
	if a.size == 0 {
		return 0, fmt.Errorf("array is empty")
	}
	sum := 0
	for i:= range a.data{
		sum+=a.data[i]
	}
	return sum,nil 
}

func (a *Array) Max()(int,int,error){
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

func (a *Array) Min() (int,int,error) {
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


func (a *Array) MaxMin() (int,int,[2]int,error) {
	if a.size == 0 {
		return 0,0,[2]int{}, fmt.Errorf("array is empty")
	}
	min, max := a.data[0] ,a.data[0]
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
	return max,min,[2]int{maxIndex,minIndex}, nil
}


func (a *Array) Insert(index int, value int) (bool, error) {
	if index < 0 || index > a.size { 
		return false, fmt.Errorf("index out of range: %d", index)
	}
	if a.size == a.capacity {
		return false, fmt.Errorf("array is full")
	}
	for i := a.size; i > index; i-- {
		a.data[i] = a.data[i-1]
	}
	a.data[index] = value
	a.size++
	return true, nil
}

func Insert(arr []int, index int, value int) ([]int, error) {
	if index < 0 || index > len(arr) { 
		return nil, fmt.Errorf("index out of range: %d", index)
	}

	arr = append(arr, 0)
	copy(arr[index+1:], arr[index:])
	arr[index] = value
	return arr, nil
}

func Max(arr []int) (int,int,error){
	if len(arr) == 0{
		return 0,0,fmt.Errorf("array is empty")
	}
	max := arr[0]
	Indexvalue := 0 
	for index := 1;index<len(arr);index++{
		if arr[index] > max {
			max = arr[index]
			Indexvalue = index
		}
	}
	return max,Indexvalue,nil 
}

func Min(arr []int) (int,int,error) {
	if len(arr) == 0 {
		return 0,0,fmt.Errorf("array is empty")
	}
	Indexvalue ,min := arr[0],0
	for index := 1; index < len(arr); index++ {
		if arr[index] < Indexvalue {
			Indexvalue = arr[index]
			min = index
		}
	}
	return Indexvalue,min,nil
	} 

func MaxMin(arr []int) (int,int,[2]int,error) {
	if len(arr) == 0 {
		return 0,0,[2]int{}, fmt.Errorf("array is empty")
	}
	max , min := arr[0] ,arr[0]
	maxIndex , minIndex := 0,0
	for i := 1; i < len(arr); i++ {
		if arr[i] > max {
			max = arr[i]
			maxIndex = i
		}
		if arr[i] < min {
			min = arr[i]
			minIndex = i
		}
	}
	return max,min,[2]int{maxIndex,minIndex}, nil
}

func Sum(arr []int) (int,error){
	if len(arr) == 0 {
		return 0, fmt.Errorf("array is empty")
	}
	sum := 0
	for i := range arr {
		sum += arr[i]
	}
	return sum,nil 
}

func ArraySlice(arr []int,start,end int) ([]int,error){
	if start < 0 || end > len(arr) || start > end {
		return nil, fmt.Errorf("index out of range: %d", start)
	}
	return arr[start:end], nil
}


func Queue(arr []int) ([]int,error){
	queue := make([]int, len(arr))
	copy(queue, arr)
	return queue,nil 
}

func Stack(arr []int) ([]int,error){
	stack := make([]int, len(arr))
	copy(stack, arr)
	return stack,nil
}

func Pop(arr []int) ([]int,error){
	if len(arr)==0{
		return nil, fmt.Errorf("array is empty")
	}
	arr  = arr[:len(arr)-1]
	return arr,nil
}

func Swap(arr []int,index1,index2 int) ([]int,error){
	if index1 < 0 || index2 > len(arr) || index1 > index2 {
		return nil, fmt.Errorf("index out of range: %d", index1)
	}
	arr[index1], arr[index2] = arr[index2], arr[index1]
	return arr, nil
}

func LSearch(arr []int,target int) (int,error){
	for index,value := range arr{ 
		if value == target{
		return index, nil
	}
	}
	return -1, fmt.Errorf("value not found")
}

func BSearch(arr []int,target int) (int,error){
	low,high := 0,len(arr)-1
	for low <= high {
		mid := (low + high) / 2
		if arr[mid] == target {
			return mid, nil
		}
		if arr[mid] < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1, fmt.Errorf("value not found")
}

func Reverse(arr []int) ([]int,error){
	for i,j :=0,len(arr)-1;i<j;i,j = i+1,j-1{
		arr[i],arr[j] = arr[j],arr[i]
	}
	return arr,nil
}


func Random(size, min, max int, noRepeat ...bool) ([]int, error) {
	if size <= 0 {
		return nil, fmt.Errorf("size must be positive: %d", size)
	}
	if min > max {
		return nil, fmt.Errorf("min must be less than or equal to max: min=%d, max=%d", min, max)
	}

	rand.Seed(time.Now().UnixNano())
	arr := make([]int, size)

	avoidDuplicates := false
	if len(noRepeat) > 0 {
		avoidDuplicates = noRepeat[0] 
	}

	if avoidDuplicates {
		if size > (max - min + 1) {
			return nil, fmt.Errorf("size is too large for unique values in the given range")
		}
		unique := make(map[int]bool)
		for i := 0; i < size; i++ {
			for {
				num := rand.Intn(max-min+1) + min
				if !unique[num] {
					arr[i] = num
					unique[num] = true
					break
				}
			}
		}
	} else {
		for i := 0; i < size; i++ {
			arr[i] = rand.Intn(max-min+1) + min
		}
	}
	return arr, nil
}
