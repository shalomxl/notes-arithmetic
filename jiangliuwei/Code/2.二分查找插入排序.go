package main

import "fmt"

func main() {
	arr := []int{20, 15, 8, 61, 79, 26, 8, 78, 56, 79}

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
		if begin >= end {
			break
		}
		mid = (begin + end) / 2
		//如果插入的数字和中间的数字相同，就放到后面，这是一个稳定的查找
		if arr[mid] > data {
			end = mid - 1
		} else {
			begin = mid + 1
		}
	}
	return begin
}
package main

import "fmt"

//func main() {
//	arr := []int{20, 15, 8, 61, 79, 26, 8, 78, 56, 79}
//
//	for i := 1; i < len(arr)-1; i++ {
//		temp:=arr[i]
//		index := BinarySearchIndex(arr[:i], i-1, arr[i])
//
//		j:=i-1
//		for {
//
//			arr[j+1]=arr[j]
//			if j==index{
//				arr[index]=temp
//				break
//			}
//			j--
//
//		}
//		fmt.Println("====>>",arr)
//
//	}
//	fmt.Println(arr)
//
//}

//func BinarySearchIndex(arr []int,end int,data int) int{
//
//	begin:=0
//	for {
//		if begin > end {
//			break
//		}
//		mid := (begin + end) / 2
//		if arr[mid]>data{
//			end=mid-1
//		}else {
//			begin=mid+1
//		}
//
//	}
//	return begin
//
//}
