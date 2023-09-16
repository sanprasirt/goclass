package main

import (
	"fmt"
	"os"

	"github.com/sanprasirt/goclass/hello"
)

func main() {
	if len(os.Args) <= 1 {
		fmt.Println(hello.Say("World"))
	} else {
		fmt.Println(hello.Say(os.Args[1]))
	}
}