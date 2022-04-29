package main

import "fmt"

// 冒泡排序算法 时间复杂度o（n*n）
// 思想：比较数组相邻的两个数字，如果左边的比右边的大，就交换位置，每轮下来大的数都会放到后面

func maopao2(arry []int) {
	for i, val := range arry {
		for j, val2 := range arry {
			if val < val2 {
				// 必须一个tmp中间变量保存上次的值
				var tmp = arry[i]
				arry[i] = arry[j]
				arry[j] = tmp
			}
		}
	}
	fmt.Println(arry)
}

func maopao(arry []int) {
	n := len(arry)
	for i := 0; i < n; i++ {
		// n-1-i 表示的是，每一轮排序之后，放到最后的值在下一轮不需要排序，所以是-i
		for j := 0; j < n-1-i; j++ {
			if arry[j] > arry[j+1] {
				//tmp := arry[j]
				//arry[j] = arry[j+1]
				//arry[j+1] = tmp
				arry[j],arry[j+1] = arry[j+1],arry[j]
			}
		}
	}
	fmt.Println(arry)
}
func main() {
	var t1 = []int{4, 3, 2, 3, 4, 5, 6, 1, 3}

	maopao(t1)

}
