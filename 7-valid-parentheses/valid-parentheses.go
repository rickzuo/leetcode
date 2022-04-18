package main

import "fmt"

//给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串 s ，判断字符串是否有效。

//有效字符串需满足：

//左括号必须用相同类型的右括号闭合。
//左括号必须以正确的顺序闭合。
//示例 1：
//
//输入：s = "()"
//输出：true

//示例 2：
//
//输入：s = "()[]{}"
//输出：true

func isValid(s string) bool {
	n := len(s)
	if n%2 == 1 {
		return false
	}
	te := map[string]string{
		")": "(",
		"]": "[",
		"}": "{",
	}
	var stack = make([]string, 0)
	var i = 0
	for i = 0; i < n; i++ {
		val := string(rune(s[i]))

		fmt.Println("data:", val, te[val])
		if _, ok := te[val]; ok {
			if len(stack) == 0 || te[val] != stack[len(stack)-1] {
				return false
			}

			// 踢出顶部的数据
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, val)
		}
	}

	return len(stack) == 0
}

func isValid2(s string) bool {
	n := len(s)
	if n%2 == 1 {
		return false

	}

	pairs := map[string]string{
		")": "(",
		"}": "{",
		"]": "[",
	}
	var stack = make([]string,0)

	for i := 0; i < n; i++ {
		val := string(s[i])
		fmt.Println(val)
		// 是否是左括号，如果是左括号就push，否则是右括号就需要判断能否消除
		if val == "(" || val == "[" || val == "{"{
			stack = append(stack, val)
		}else {
			// 如果栈内为空，现有的右括号没有匹配对象 返回false
			if len(stack) == 0 {
				return false
			}
			// 如果栈顶的元素与当前的括号相匹配的元素不相等，则返回false
			if pairs[val] != stack[len(stack)-1]{
				return false
			}
			// 匹配中了，栈顶元素出栈，相当于消除了。
			stack = stack[:len(stack)-1]
		}
	}
	return len(stack) == 0
}

func main() {
	var t = "{)"
	ret := isValid2(t)
	fmt.Println(ret)
}
