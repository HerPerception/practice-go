package main

import "fmt"
import "os"

func main(){
	if len(os.Args) != 2 {
		return
	}
	s := os.Args[1]
	if s == " " || s == "" {
		return
	}
	word := []string{}
	i := 0
	n := 0
	for index, ch := range s {
		if index == len(s)-1 && ch == ' ' {
			index--
		} else if ch != ' ' {
			i = index
			index--
		} else if ch == ' ' {
			n = index + 1
		}
	}
	word = append(word, s[n:i+1])
	text := ""
	for _, ch := range word {
		text += string(ch)
	}
	fmt.Println(text)
} fmt.Printf("%q\n", itoa(12345))