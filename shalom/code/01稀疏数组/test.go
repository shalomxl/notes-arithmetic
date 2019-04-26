package main

import "fmt"

func main() {

	buf := make([][]int, 10, 10)

	for _, v := range buf {
		v = make([]int,11)
		for _, val := range v {
			fmt.Print(val, " ")
		}
		fmt.Println()
	}

}
