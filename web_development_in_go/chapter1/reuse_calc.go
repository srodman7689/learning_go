package main

import (
	"fmt"
	"github.com/srodman7689/learning_go/web_development_in_go/chapter1/calc"
)

func main() {
	var x, y int = 10, 5
	fmt.Println(calc.Add(x, y))
	fmt.Println(calc.Subtract(x, y))
}
