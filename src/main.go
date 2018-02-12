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
	"runtime"
	"time"
)

func GetTid() int {
	return int(C.gettid())
}

func PrintTid() {
	fmt.Print("tid:", GetTid(), "\n")
}

func IntSleep() {
	PrintTid()
	time.Sleep(5 * time.Second)
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	fmt.Print("app run tid:", GetTid(), "\n")
	for i := 1; i < 10; i++ {
		go func() {
			for j := 1; j < 5; j++ {
				PrintTid()
			}
		}()
	}
	time.Sleep(6000 * time.Second)
}
