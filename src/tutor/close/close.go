package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	go func() {
		time.Sleep(time.Second)
		close(ch)
	}()

	for e := range ch {
		fmt.Println(e)
	}
	/*v, ok := <-ch
	fmt.Println(<-ch)
	fmt.Println(<-ch)*/
}
