package main

import "fmt"

func isValid3(s string) bool {
	var n = len(s)
	if n % 2 != 0 {
		return false
	}

	var stack = make([]string,0)
	var m = map[string]string{
		"}":"{",
		"]":"[",
		")":"(",
	}

	for i := 0 ;i < n;i ++ {
		var val = string(rune(s[i]))
		if left,ok := m[val];ok{
			if len(stack) == 0 {
				return false
			}
			if left != stack[len(stack) - 1]{
				return false
			}
			stack = stack[:len(stack) -1 ]

		}else{
			stack  = append(stack,val)
		}
	}
	return len(stack) == 0
}


func main() {
	var t = "{{[()]}}"
	res := isValid3(t)
	fmt.Println(res)
}
