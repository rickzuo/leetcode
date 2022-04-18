//两个协程交替打印1-100的奇数偶数
package main

import (
	"fmt"
	"time"
)
func print(i int){
	fmt.Println(i)
}

func main(){

	var data = make([]int,0)
	var i = 0
	for  i = 0;i<100;i++{
		data = append(data,i+1)
	}
	var ch = make(chan int,0)
	fmt.Println(data)

	var ch2 = make(chan int ,0)
	//go func(){
	//	<- ch2
	//}()
	//ch2 <- 100

	ch2 <- 100

	c2,ok := <- ch
	if ok{
		fmt.Println("c2:",c2)
	}

	fmt.Println()


	for _,j := range data {
		if j % 2 == 0 {
			go func(chan int) {
				ch <- j
				fmt.Println(j)
			}(ch)
		}else{
			go func(chan int) {
				ch <- j
				fmt.Println(j)
			}(ch)
		}
		<- ch
	}
	time.Sleep(2)
}