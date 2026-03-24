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
	for index := 0; index < len(arg); index++ { 
		if arg[index] != ' '{
			break
		}
		start = index
    }last := text[index-1]
	for index := len(arg)-1; index >= 0; index-- { 
		if arg[index] != ' '{
			break
		}
			end = index
    }
	text = append(text, arg[start + 1:end])

	for index := 0; index < len(text); index++ {
	last := index-1
		if text[index] == " " && text[last] == " " {
			continue
		} else if text[index] == " " && text[last] != " " {
			word += text[index]
			last = index
		} else {
			word += text[index] 
		}
	}
	fmt.Printf("%q\n",word)
}