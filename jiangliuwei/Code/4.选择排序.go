package main

import "fmt"

func main() {

	arr2 := []int{27, 15, 58, 89, 45, 25}

	newarr := selectionSort(arr2)
	fmt.Println(newarr)

}

func selectionSort(arr []int) []int {

	for i := 0; i < len(arr); i++ {
		Temp := arr[i]
		var Index int
		for j := i; j < len(arr); j++ {
			if arr[j] < Temp {
				Temp = arr[j]
				Index = j
			}
		}
		//Index==0表示此时他就是最小的
		if Index != 0 {
			arr[i], arr[Index] = arr[Index], arr[i]
		}
		fmt.Println(arr)
	}
	return arr
}
