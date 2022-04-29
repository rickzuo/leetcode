package main

import "fmt"

// 选择排序 - 不稳定的排序，相同的元素经过排序之后，相同元素的相对位置发现变化了，序列arr = [5 8 5 2 9]，第一遍选择第1个元素5会和2交换，那么两个5的相对前后顺序就被破坏了，如果是冒泡2个5的前后位置不会发生改变
// 思想：第一次选择最小的数，放到数组前面，下次又去选择最小的数，并放到已排序的数组末尾，注意是已排序的数组。
func xuanzhe(arr1 []int) {
	var n = len(arr1)
	for i := 0; i < n; i++ {
		// min是用来记录下标值，假定第一个下标i是第一个min
		var min = i
		for j := i + 1; j < n; j++ {
			// 比较最小的min和下一个数组
			if arr1[j] < arr1[min] {
				// 如果下一个数组比min还小，那就min记录最小的数j
				min = j
			}
		}
		// 最外层i和min交互位置
		arr1[i], arr1[min] = arr1[min], arr1[i]
	}
	fmt.Println(arr1)
}

func xuanzhe2(arr1 []int) {

	var n = len(arr1)
	for i := 0; i < n; i++ {
		var min = i
		for j := i + 1; j < n; j++ {
			if arr1[min] > arr1[j] {
				min = j
			}
		}
		arr1[i], arr1[min] = arr1[min], arr1[i]

	}
	fmt.Println(arr1)
}
func main() {
	var arr1 = []int{1, 3, 9, 3, 1, 4, 5, 8, 12, 11}
	xuanzhe2(arr1)
}
