package main

import "fmt"

//给你两个按 非递减顺序 排列的整数数组nums1 和 nums2，另有两个整数 m 和 n ，分别表示 nums1 和 nums2 中的元素数目。
//请你 合并 nums2 到 nums1 中，使合并后的数组同样按 非递减顺序 排列。
//注意：最终，合并后数组不应由函数返回，而是存储在数组 nums1 中。为了应对这种情况，nums1 的初始长度为 m + n，其中前 m 个元素表示应合并的元素，后 n 个元素为 0 ，应忽略。nums2 的长度为 n 。

func merge(nums1 []int32, nums2 []int32, m, n int32) []int32 {

	if n == 0 {
		return nums1
	}
	//先定义3个下标
	var k = m - 1     //不为0的num1数组下标
	var j = n - 1     // num2数组下标
	var p = m + n - 1 //整个num1数组的长度下标

	// 循环条件是当遍历完num1 不为0的长度和num2的长度
	for k >= 0 && j >= 0 {
		// 比较num1和num2 数组最后的值，如果num1最后不为0的值小于num2的最后的值
		if nums1[k] < nums2[j] {
			// 则将num1最后的值替换为num2的值
			nums1[p] = nums2[j]
			j-- // 替换num1数组后面的值，需要将数组2的下标前移一位
		} else {
			// 说明num1值大于num2的值，则将num1数组后面的值替换 num1本身大于num2的值
			nums1[p] = nums1[k]
			k-- // 并且将数组1的下标前移
		}
		// 每轮下来都需要将整个数组p的下标前移动一位
		p--
	}
	// 如果上面比较之后，发现num2的数组还有剩余，说明这些数字都
	for j >= 0 {
		nums1[p] = nums2[j]
		j--
		p--
	}

	return nums1
}

func main() {

	var a = []int32{4, 5, 6, 0, 0, 0}
	var b = []int32{1, 2, 3}
	c := merge(a, b, 3, 3)

	fmt.Println("c:", c)

}
