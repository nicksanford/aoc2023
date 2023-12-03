package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprint(os.Stderr, "usage: <cmd> file_to_parse")
		os.Exit(1)
	}

	b, err := os.ReadFile(os.Args[1])
	if err != nil {
		fmt.Fprint(os.Stderr, err.Error())
		os.Exit(1)
	}

	input := string(b)
	var sum int
	for _, c := range input {
		switch c {
		case '(':
			sum++
		case ')':
			sum--
		default:
		}
	}

	fmt.Print(sum)
}
