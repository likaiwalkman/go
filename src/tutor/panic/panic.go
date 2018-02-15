package main

import "fmt"

func r() {
	if e := recover(); e != nil {
		fmt.Println(e)
	}
}

func main() {

	defer r()
	defer func() {fmt.Println("1st defer")}()
	defer func() {fmt.Println("2nd defer")}()
	//defer r()
	defer func() {fmt.Println("3rd defer")}()
	//defer r()
	panic("ouch")
}