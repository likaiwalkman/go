package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

func HelloServer(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "hello, world!\n")
}

func main() {
	http.HandleFunc("/hello", HelloServer)
	err := http.ListenAndServe(":8888", nil)
	fmt.Print("listener ends")
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
	for {
		time.Sleep(5 * time.Second)
	}

}
