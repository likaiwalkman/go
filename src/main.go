package main

/*
#include <unistd.h>
#include <sys/syscall.h>
#include <stdlib.h>
#include <sys/types.h>
int gettid(){
	//SYS_gettid for linux
    return syscall(SYS_thread_selfid); //for Mac
}
*/
import "C"
import (
	"fmt"
	//"runtime"
	"os"
	"strconv"
	"time"
)

func GetTid() int {
	return int(C.gettid())
}

func PrintTid() {
	var now = time.Now()
	fmt.Printf("now:%d:%d:%d, print tid:%d\n", now.Hour(), now.Minute(), now.Second(), GetTid())
}

func IntSleep() {
	PrintTid()
	time.Sleep(1 * time.Second)
}

func main() {
	//runtime.GOMAXPROCS(runtime.NumCPU())
	fmt.Println("app run tid:", GetTid(), "\n")

	var ch = make(chan string)

	go func() {
		for i := 0; i < 5; i++ {
			go func(i int) {
				for j := 0; j < 3; j++ {
					IntSleep()
					ch <- strconv.Itoa(i) + "," + strconv.Itoa(j)
				}
			}(i)
		}
	}()

	//go func() {
	var timer = time.NewTimer(20 * time.Second)
	for {
		select {
		case v := <-ch:
			fmt.Printf("main-loop receive:%s\n", v)
		case <-timer.C:
			fmt.Println("exit")
			os.Exit(0)
		}

	}
	//}()

}
