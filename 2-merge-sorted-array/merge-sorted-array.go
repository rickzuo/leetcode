package main

import "fmt"

func merge(nums1 []int, m int, nums2 []int, n int) {

	var k = m -1
	var j = n -1
	var p = m + n -1

	for k >= 0 && j >=0 {
		if nums1[k] <= nums2[j]{
			nums1[p] = nums2[j]
			j--
		}else{
			nums1[p] = nums1[k]
			k--
		}
		p--
	}
	for j >= 0{
		nums1[p] = nums2[j]
		j--
		p--
	}

	fmt.Println("nums1:",nums1)
}

func main() {
	var num1 = []int{1, 2, 3, 0, 0, 0}
	var num2 = []int{4, 5, 6}
	merge(num1, 3, num2, 3)
}
