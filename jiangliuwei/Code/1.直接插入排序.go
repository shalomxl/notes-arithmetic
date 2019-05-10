package main

import "fmt"

func main() {

	var arr = []int{20, 19, 5, 13, 89, 56, 23}
	NewArr := InsertionSort(arr)
	fmt.Println(NewArr)

}
func InsertionSort(arr []int) []int {

	for i := 1; i < len(arr); i++ {

		insertval := arr[i]
		insertIndex := i - 1
		for insertIndex >= 0 && arr[insertIndex] < insertval {
			arr[insertIndex+1] = arr[insertIndex]
			insertIndex--
		}
		if insertIndex+1 != i {
			arr[insertIndex+1] = insertval
		}
	}
	return arr
}
