package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	fmt.Println("ハロー")

	b, err := ioutil.ReadFile("ハロー.txt")
	if nil != err {
		fmt.Println("read file failed: ", err.Error())
	}

	fmt.Println(b)

	s := string(b)

	fmt.Println(s)
}
