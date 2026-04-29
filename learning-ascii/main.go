package main

import (
	"fmt"
	"os"
)

func main() {
	result, err := LoadBanner("standard.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(Generate(os.Args[1], result))
}
