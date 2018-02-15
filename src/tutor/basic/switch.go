package main

import (
	"runtime"
	"fmt"
)

func main(){
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("Mac.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.", os)
	}

}
