package main

import "fmt"

/*
给定两个大小为 m 和 n 的有序数组 nums1 和 nums2。

请你找出这两个有序数组的中位数，并且要求算法的时间复杂度为 O(log(m + n))。

你可以假设 nums1 和 nums2 不会同时为空。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/median-of-two-sorted-arrays
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */

//  第一种方法，将两个有序数组合并成一个有序数组，然后取出中位数
func findMedianSortedArrays01(nums1 []int, nums2 []int) float64 {
	//m := len(nums1)
	//n := len(nums2)
	//	将两个有序的数组合并到nums中，并且保证nums也是有序的
	nums1 = mergeArrs(nums1, nums2)
	fmt.Println(nums1)

	l := len(nums1)
	var res float64
	if l%2 == 0 {
		res = (float64(nums1[l/2]) + float64(nums1[l/2-1])) / 2
	} else {
		res = float64(nums1[l/2])
	}

	return res
}

func mergeArrs(nums1 []int, nums2 []int) []int {
	m := len(nums1)
	n := len(nums2)

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
	copy(nums1, nums2[:p2+1])
	return nums1
}

/*
	时间复杂度：O(m+n)
	空间复杂度：O(1)
 */

//	第二种方法：
func findMedianSortedArrays02(nums1 []int, nums2 []int) float64 {
	m := len(nums1)
	n := len(nums2)

	if m > n {
		m, n = n, m
		nums1, nums2 = nums2, nums1
	}
	fmt.Println(m,n)
	fmt.Println(nums1)
	fmt.Println(nums2)

	//	i的取值范围，i有m+1种可能
	iMin := 0
	iMax := m

	for iMin <= iMax {
		i := (iMin + iMax) / 2
		j := (m+n+1)/2 - i

		//	当i==iMax的时候
		if i > iMin && nums1[i-1] > nums2[j]  {
			//	说明i大了，所以这个时候要判断i是否还有减小的空间,加上 && i>iMin 来确保i还有减少空间，没有到达临界值
			iMax = i - 1
		} else if i < iMax && nums2[j-1] > nums1[i]  {
			iMin = i + 1
		} else {
			//	走到这里，就已经可以确定中位数了，剩下的问题就是，如果元素个数是偶数，就要取出最中间两个数，然后求平均值
			//	可能出现的临界值：i == 0 ; j == 0 ; i == m ; j == n
			//		其中，j == 0 和 j == n 都仅在两个数组长度一样的时候才会出现
			//	无论是否出现临界值，都可以判断出中位数
			var maxLeft int
			if i == 0 {
				maxLeft = nums2[j-1]
			} else if j == 0 {
				//	只有当 m==n 的时候，j才有可能为0
				maxLeft = nums1[i-1]
			}else {
				if nums1[i-1] > nums2[j-1] {
					maxLeft = nums1[i-1]
				} else {
					maxLeft = nums2[j-1]
				}
			}


			if (m+n)%2 == 1 {
				return float64(maxLeft)
			}

			var minRight int
			if i == m {
				minRight = nums2[j]
			} else if j == n {
				minRight = nums1[i]
			}else {
				if nums1[i] < nums2[j] {
					minRight = nums1[i]
				} else {
					minRight = nums2[j]
				}
			}

			return (float64(minRight) + float64(maxLeft)) / 2.0
		}
	}
	return 0
}

func main() {
	nums1 := []int{1, 2, 3, 4, 5, 6}
	nums2 := make([]int, 0)

	for i := 0; i < 10; i++ {
		nums2 = append(nums2, i+10)
	}

	res := findMedianSortedArrays02(nums1, nums2)
	fmt.Println(res)
}
