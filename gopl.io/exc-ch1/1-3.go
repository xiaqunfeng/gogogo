package main

import (
	"fmt"
	"os"
	"strings"
)

func echo2() {
	s, sep := "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}

	fmt.Println(s)
}

func echo3() {
	fmt.Println(strings.Join(os.Args[1:], " "))
}

func main() {
	echo2()
	echo3()
}
