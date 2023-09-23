package main

import (
	"fmt"

	"github.com/samk-dev/toolkit"
)

func main() {
	var tools toolkit.Tools

	s := tools.RandomStringGenerator(10)

	fmt.Println("Random string:", s)
}
