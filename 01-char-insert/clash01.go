package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("c = ")
	c, _ := fmt.Scan()
	fmt.Println("l = ")
	lin, _ := reader.ReadString('\n')
	l, lerr := strconv.ParseInt(lin, 10, 32)

	if (lerr != nil) {
		fmt.Println(lerr)
	}

    ans := ""

	if l % 2 != 0 && l > 2 {
		pad := strings.Repeat("*", int(l) / 2)
		ans = pad + c + pad
	} else {
		ans = "CAN'T"
	}

	fmt.Println(ans)
	fmt.Println(c)
	fmt.Println(l)
	fmt.Println(lin)
}
