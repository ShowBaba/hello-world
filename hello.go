package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	who := "World!"
	if len(os.Args) > 1 { /* os.Args[0] is "hello" or "hello.exe" */
		fmt.Println(os.Args)
		who = strings.Join(os.Args[1:], " ")
	}
	fmt.Println("Hello", who)
}
