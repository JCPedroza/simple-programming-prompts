package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	errorSize = "Expected two arguments"
	errorType = "Expected types: arg1 char, arg2 int"
	errorSpec = "CAN'T"
	fmtOut = "%s%s%s\n"
	padChar = "*"
)

func main() {
	mid, size, err := parseInput()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if size % 2 == 0 || size < 3 {
		fmt.Println(errorSpec)
	} else {
		pad := strings.Repeat(padChar, size / 2)
		fmt.Printf(fmtOut, pad, mid, pad)
	}
}

func parseInput() (string, int, error) {
	if len(os.Args) != 3 {
		return "", 0, errors.New(errorSize)
	}

	mid := os.Args[1]
	size, err := strconv.Atoi(os.Args[2])
	if len(mid) != 1 || err != nil {
		return "", 0, errors.New(errorType)
	}

	return mid, size, nil
}
