package main

import (
	"fmt"
	"log"
	"strings"
)

// 给定字符串，按照规定n的长度拆分多个字符

func split(str string, n int) (ret []string, err error) {
	data := strings.Fields(strings.TrimSpace(str))
	res := ""
	for _, d := range data {
		if len(res) < n {
			res = res + " " + d
		}else{
			ret = append(ret, res)
			res =  d
		}
	}
	if res != ""{
		ret = append(ret, res)
	}
	return
}

func main() {
	var t = " i want  not go  work , must be strong"
	res,err := split(t,10)
	if err != nil {
		log.Fatal(err)
	}
	for _,r := range res{
		fmt.Println(r)
	}
}
