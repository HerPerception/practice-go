package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	data, err := os.ReadFile("standard.txt")
	if err != nil {
		return
	}

	lines := strings.Split(string(data), "\n")

	str := "H"

	for r := 0; r < 8; r++ { // FIX: must be < 8, not <= 8
		row := []string{}

		for _, c := range str {
			index := int(c-32)*8 + r
			end := index + 7
			if index >= 0 && index < len(lines) {
				row = append(row, lines[index:end]...)
			}
		}

		fmt.Println(strings.Join(row, ""))
	}
}
