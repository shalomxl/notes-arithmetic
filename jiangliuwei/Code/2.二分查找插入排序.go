package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7}

	//将arr分成无序区和有序区，初始的有序区只有一个元素
	//将i-1分为有序区，将i~len（arr）-1分为无序区
	for i := 1; i < len(arr); i++ {
		index := BinarySearchIndex(arr, i, arr[i])
		temp := arr[i]
		//这一步和插入排序的逻辑类似
		//最坏的情况为1+2+3+...N次
		for j := i - 1; j >= index; j-- {
			arr[j+1] = arr[j]
			if j == index {
				arr[index] = temp
			}
		}
	}
	fmt.Println(arr)
}

//这一步通过二分查找，查找到对应的插入的下表
func BinarySearchIndex(arr []int, maxIndex int, data int) int {
	begin := 0
	end := maxIndex
	var mid int
	for {
		if begin == end {
			break
		}
		mid = (begin + end) / 2
		//如果插入的数字和中间的数字相同，就放到后面，这是一个稳定的查找
		if arr[mid] > data {
			end--
		} else {
			begin++
		}
	}
	return begin
}
