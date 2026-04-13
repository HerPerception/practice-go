package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Include text.")
		return
	}
	str := os.Args[1]
	data, err := os.ReadFile("standard.txt")
	if err != nil {
		return
	}

	lines := strings.Split(string(data), "\n")

	for r := 0; r < 8; r++ { // FIX: must be < 8, not <= 8
		row := []string{}

		for c := range str {
			index := int(str[c]-32)*9 + r
			if index >= 0 && index < len(lines) {
				row = append(row, lines[index])
			}
		}

		fmt.Println(strings.Join(row, ""))
	}
}
