package main

import (
	"fmt"
	"os"
	"strconv"
)

func fibTool(num int) int {
	if num == 0 || num == 1 {
		return num
	}
	return fibTool(num - 1) + fibTool(num - 2)
}

func main() {
	if len(os.Args) < 2 {
		fmt.Fprint(os.Stderr, "do nothing\n")
		os.Exit(1)
	}
	i, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "argument must be integer: %s\n", err)
		os.Exit(1)
	}
	fmt.Printf("%d", fibTool(i))
}
