package main

import(
	"fmt"
	"sync"
	"time"
)
func main(){
	ch1 := make(chan struct{} ,1)
	ch2 := make(chan struct{} ,1)

	wait := sync.WaitGroup{}

	go func(group *sync.WaitGroup){
		defer group.Done()
		var i = 1
		for{
			select {
				case <- ch2:
					fmt.Print(i)
					i++
					fmt.Print(i)
					i++
					ch1 <- struct{}{}
					time.Sleep(time.Second * 1)
			}
		}
	}(&wait)
	wait.Add(1)

	go func(group *sync.WaitGroup){
		defer group.Done()
		var j = 'a'
		for{
			select {
			case <- ch1:
				// 注意是string(jf)
				fmt.Print(string(j))
				j++
				fmt.Print(string(j))
				j++
				ch2 <- struct{}{}
				time.Sleep(time.Second * 1)
			}
		}
	}(&wait)

	wait.Add(1)
	//ch1 <- struct{}{}
	ch2 <- struct{}{}

	wait.Wait()
	close(ch1)
	close(ch2)

}