package main

import "fmt"

func merge2(nums1, nums2 []int) {
	var m = len(nums1) - 1
	var n = len(nums2) - 1
	var p = len(nums1) + len(nums2) - 1

	var tmp = make([]int, p+1)
	fmt.Println(tmp,p)
	for m > 0 && n > 0 {
		if nums1[m] <= nums2[n] {
			tmp[p] = nums2[n]
			n--
		} else {
			tmp[p] = nums1[m]
			m--
		}
		p--
	}
	fmt.Println(m,n,tmp)
	for n >= 0 {
		tmp[p] = nums2[n]
		n--
		p--
	}
	for m >= 0 {
		fmt.Println("p:",p,m)
		tmp[p] = nums1[m]
		m--
		p--
	}

	fmt.Println(tmp)
}

func main() {
	var num1 = []int{1, 2, 3}
	var num2 = []int{4, 5, 6}
	merge2(num1, num2)
}
