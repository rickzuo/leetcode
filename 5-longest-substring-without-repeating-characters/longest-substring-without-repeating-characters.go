package main

import (
	"bytes"
	"fmt"
)

//输入: s = "abcabcbb"
//输出: 3
//解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。

//输入: s = "bbbbb"
//输出: 1
//解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。

//输入: s = "pwwkew"
//输出: 3
//解释: 因为无重复字符的最长子串是 "wke"，所以其长度为 3。
//请注意，你的答案必须是 子串 的长度，"pwke" 是一个子序列，不是子串。

func Index(s string, index int) string {
	runes := bytes.Runes([]byte(s))
	for i, data := range runes {
		if i == index {
			return string(data)
		}
	}
	return ""
}

//lengthOfLongestSubstring 给定一个字符串 s ，请你找出其中不含有重复字符的 最长子串 的长度。
func lengthOfLongestSubstring2(s string) int {
	var maxLen = 0
	var start = 0
	var end = 0
	var window = make(map[rune]int, 0)
	sv := []rune(s)
	for end < len(s) {
		val := sv[end]
		lastIndex, ok := window[val]
		if ok && lastIndex >= start {
			delete(window, val)
			start = lastIndex + 1
		}
		if maxLen < end-start+1 {
			maxLen = end - start + 1
		}
		window[val] = end
		end++
		fmt.Println(start, end, val, maxLen)
	}
	return maxLen

}

//lengthOfLongestSubstring 给定一个字符串 s ，请你找出其中不含有重复字符的 最长子串 的长度。
func lengthOfLongestSubstring3(s string) int {
	var start = 0
	var maxLen = 0
	var set = make(map[rune]int,0)
	for end,val := range []rune(s){
		if index,ok := set[val];ok && index >= start{
			start = index + 1
		}
		set[val] = end
		if maxLen < end - start + 1{
			maxLen = end - start + 1
		}

	}
	return maxLen

}

//lengthOfLongestSubstring 给定一个字符串 s ，请你找出其中不含有重复字符的 最长子串 的长度。
func lengthOfLongestSubstring(s string) int {
	// 定义好滑动窗口5要素；前后指针，无重复的队列，最大长度，当前长度
	var maxLen = 0
	var start = 0
	var curLen = 0
	var window = make(map[rune]int, 0)
	sv := []rune(s)

	for end, val := range sv {
		if lastIndex, ok := window[val]; ok && lastIndex >= start {
			start = lastIndex + 1
		}
		curLen = end - start + 1
		window[val] = end
		if curLen > maxLen {
			maxLen = curLen
		}
	}
	return maxLen

}

func main() {
	//var t = "pwwkew"
	//var t = "abcabcbb"
	//var t = "bbbbb"
	var t = "abba"

	ret := lengthOfLongestSubstring3(t)
	fmt.Println(ret)
}
