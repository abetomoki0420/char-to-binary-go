package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Please input args.")
		return
	}

	argv := os.Args[1]

	r := []rune(argv)
	for i := 0; i < len(r); i++ {
		s := string(r[i])
		byteSize := len([]byte(s))

		fmt.Printf("%d文字目 %s\n", i+1, string(r[i]))
		for i := 0; i < byteSize; i++ {
			fmt.Printf("[%d]:%b\n", i, s[i])
		}

		fmt.Println("")
	}
}
