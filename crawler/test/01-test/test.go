package main

import (
	"fmt"
	"sync"
)

var Lock sync.Mutex
func main()  {
	num := 0
	ch := make(chan int)
	for i := 0; i < 100000; i++{
		go func() {
			Lock.Lock()
			num++
			Lock.Unlock()
			if num == 100000 {
				ch <- 1
			}
		}()
	}
	<- ch

	fmt.Println(num)
}
