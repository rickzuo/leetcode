package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

//输入：l1 = [1,2,4], l2 = [1,3,4]
//输出：[1,1,2,3,4,4]

//输入：l1 = [], l2 = [0]
//输出：[0]

//输入：l1 = [], l2 = []
//输出：[]

//mergeTwoLists 将两个升序链表合并为一个新的 升序 链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {

	var retListNode = &ListNode{}
	var cur = retListNode

	// 同时遍历2个链表
	for list1!= nil && list2 != nil {

		if list1.Val <= list2.Val {
			// cur指针的下一个指向当前的小的数
			cur.Next = list1
			// cur指针后移
			cur = cur.Next
			// 链表1指针后移
			list1 = list1.Next
		} else {
			cur.Next = list2
			// 指针后移
			cur = cur.Next
			// 链表2指针后移
			list2 = list2.Next
		}
	}
	// 遍历完一边后，把剩余的直接拼接到cur
	if list1 == nil {
		cur.Next = list2
	}else{
		cur.Next = list1
	}
	// 因为头部开始是0，返回的是从下一个开始
	return retListNode.Next

}

func main() {
	var list1 = &ListNode{
		Val: 4,
		Next: &ListNode{
			Val: 6,
			Next: &ListNode{
				Val: 7,
				Next: &ListNode{
					Val: 10,
					Next: nil,
				},
			},
		},
	}
	var list2 = &ListNode{
		Val: 3,
		Next: &ListNode{
			Val: 5,
			Next: &ListNode{
				Val: 7,
				Next: &ListNode{
					Val: 9,
					Next: nil,
				},
			},
		},
	}
	ret := mergeTwoLists(list1, list2)
	for ret != nil {
		fmt.Println(ret.Val)
		ret = ret.Next
	}
}
