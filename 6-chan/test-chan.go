package main

import (
	"fmt"
	"time"
)

func t1(){
	var ch = make(chan int, 0)
	// 无缓存chan 必须起一个goroutine 负责读，然后再写
	go func() {
		if data, ok := <-ch; ok {
			fmt.Println(data)
		}
	}()
	ch <- 100
	close(ch)
}

func t2(){
	// 注意这里是有缓存的chan，如果把1改为0会painc，因为没有另外一个grouting在读ch
	var ch = make(chan int, 1)
	ch <- 101
	go func() {
		select {
		case data:= <- ch:
			fmt.Println(data)
		}
	}()
	close(ch)
}
func main() {
	t1()
	t2()
	time.Sleep(20)
}

