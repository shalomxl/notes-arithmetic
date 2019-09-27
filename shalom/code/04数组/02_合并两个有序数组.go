package main

import "fmt"

/*
	给定两个有序整数数组 nums1 和 nums2，合并成为一个有序数组。
 */
func merge(nums1 []int,nums2 []int) []int {
	m:=len(nums1)
	n:=len(nums2)

	space := make([]int, n)
	nums1 = append(nums1, space...)

	p1 := m - 1
	p2 := n - 1
	p := m + n - 1

	for p1 >= 0 && p2 >= 0 {
		if nums1[p1] >= nums2[p2] {
			nums1[p] = nums1[p1]
			p1--
		} else {
			nums1[p] = nums2[p2]
			p2--
		}
		p--
	}

	//	将 nums2 中剩余的元素拷贝到 nums1 中
	copy(nums1,nums2[:p2+1])
	return nums1
}
/*
总结：
	1. 在操作数组时，可以自定义下标，作为元素的指针，并且灵活运用这些指针，已达到想要的效果；
	2. 从时间复杂度方面考虑：应该利用好原有的顺序，可以减少排序的步骤；
	3. 从空间复杂度方面考虑：已经存在的空间能否利用？例：上例中可以利用nums1的空间，所以可以将nums2合并到nums1中，不需要第三个数组；

	时间复杂度：O(m+n)
	空间复杂度：O(1)  因为m+2n是原本就需要的，不是因为算法额外增加的
 */

func main() {
	nums1 := []int{1, 2, 3, 4, 5}
	nums2 := make([]int, 0)


	for i := 0; i < 10; i++ {
		nums2 = append(nums2, i+10)
	}
	fmt.Println(nums1)
	fmt.Println(nums2)

	res := merge(nums1, nums2)
	fmt.Println(res)
}
