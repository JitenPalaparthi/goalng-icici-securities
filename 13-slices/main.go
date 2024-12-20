package main

import (
	"errors"
	"fmt"
)

func main() {

	//slice1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	slice1 := make([]int, 10, 10)
	for i := range slice1 {
		slice1[i] = i + 1
	}
	slice8 := slice1

	// slice2 := slice1[:]   // 0th to the end
	// slice3 := slice1[:4]  // 0th to the 4th
	// slice4 := slice1[4:]  // 4th to the end
	// slice5 := slice1[4:8] // 4th to the 8th
	// Printheader(slice1, "slice1")
	// Printheader(slice2, "slice2")
	// Printheader(slice3, "slice3")
	// Printheader(slice4, "slice4")
	// Printheader(slice5, "slice5")

	// arr1 := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	// slice6 := arr1[4:]
	// slice7 := arr1[4:8]

	// Printheader(slice6, "slice6")
	// Printheader(slice7, "slice7")

	Printheader(slice8, "slice8")
	slice8 = SquareAndAdd(slice8, 11, 12, 13, 14, 15)
	Printheader(slice8, "slice8")

	if slice8, err := RemoveElem(slice8, 5); err != nil {
		println(err.Error())
	} else {
		Printheader(slice8, "slice8")
	}

}

func Printheader(slice []int, name string) error {
	if slice == nil {
		return errors.New("nil slice")
	}
	fmt.Printf("Slice:%s Address:%p Values: %v ptr:%p Len: %d Cap: %d\n", name, &slice, slice, &slice[0], len(slice), cap(slice))
	return nil
}

func SquareAndAdd(slice []int, nums ...int) []int {
	for _, v := range nums {
		slice = append(slice, v)
	}
	for i, v := range slice {
		slice[i] = v * v
	}
	return slice
}

// 1 2 3 4 5 6 7 8 9 10

// slice[:index] = 1 2 3 4
// slice[index+1:]= 6 7 8 9 10
// 1 2 3 4 6 7 8 9 10
func RemoveElem(slice []int, index int) ([]int, error) {
	if slice == nil {
		return nil, errors.New("nil slice")
	}
	if index >= len(slice)-1 {
		return nil, errors.New("invalid index")
	}
	slice = append(slice[:index], slice[index+1:]...)
	return slice, nil
}
