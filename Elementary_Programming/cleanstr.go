package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println()
		return
	}
	arg := os.Args[1]
	text := []string{}
	start := 0
	end := 0
	word := ""
	for index, ch := range arg {
		 if (index >= 0 && index < len(arg)/2) && ch == ' ' {
           index++
		 } else if (index >= 0 && index < len(arg)/2) && ch != ' '{
			start = index
		}
		if index == len(arg)-1 && ch == ' ' {
			index--
		} else if ch != ' ' {
			end = index
		}
	}
	text = append(text, arg[start:end])
	for _, ch := range text {
		word += string(ch)
	}
	fmt.Printf("%q\n", word)
}