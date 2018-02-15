package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int, 3)
	ch <- 1
	ch <- 2
	ch <- 3
	go func() {
		time.Sleep(3*time.Second)
		close(ch)
	}()

	v, open := <-ch
	fmt.Println(v,open)
	
	var b bool
	for{

		select {
		case v, open := <-ch:
			b = open
			if open {
				fmt.Println(v, open)
				time.Sleep(1*time.Second)
			}
		default:
			fmt.Println("default")
			time.Sleep(1*time.Second)
		}
		if b == false {
			break
		}
	}

	/*v, ok := <-ch
	fmt.Println(<-ch)
	fmt.Println(<-ch)*/
}
