package main

import (
	"fmt"
	"github.com/srodman7689/learning_go/web_development_in_go/chapter2/strcon"
)

func main() {
	s := strcon.SwapCase("Gopher")
	fmt.Println("Converted string is :", s)
}
