package main

import "fmt"

func main() {

	arr1 := []int{20, 15, 18, 79, 35, 24, 18}

	newArr := ShellSort(arr1)
	fmt.Println(newArr)

}

func ShellSort(arr []int) []int {

	d := len(arr) / 2
	for d > 0 {
		for i := 0; i < d; i++ {
			for j := i + d; j < len(arr); j++ {
				Temp := arr[j]
				k := j - d
				for k >= 0 && arr[k] > Temp {
					arr[k+d] = arr[k]
					k -= d
				}
				arr[k+d] = Temp
			}
		}
		d = d / 2
	}
	return arr
}
