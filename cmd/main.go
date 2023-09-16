package main

import (
	"fmt"
	"os"

	"github.com/sanprasirt/goclass"
)

func main() {
	fmt.Println(goclass.Say(os.Args[1:]))
}