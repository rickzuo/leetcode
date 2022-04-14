package main

import "fmt"

//输入：nums = [2,7,11,15], target = 9
//输出：[0,1]
//解释：因为 nums[0] + nums[1] == 9 ，返回 [0, 1] 。

func twoSum(nums []int,target int) []int {
	// 需要一个map保存遍历的数组值和下标
	var m = make(map[int]int,0)
	for aIndex,a := range nums{
		// 计算得出要找到的对值数值
		left := target - a
		// 到map中找，是否存在另外值的
		if index,ok := m[left];ok{
			// 如果存在，则返回下标值，注意返回顺序，后入的在前
			return []int{index,aIndex}
		}
		// 记录所有数组的值和下标
		m[a] = aIndex
	}
	return nil
}

func main() {
	var nums = []int{2, 7, 11, 15}
	var target = 13
	ret := twoSum(nums,target)
	fmt.Println(ret)
}
