package main

import "strings"

func RenderLines(text string, banner map[rune][]string) []string {
	var slice []string
	for r := 0; r < 8; r++ {
		var builder strings.Builder
		for _, ch := range text {
			builder.WriteString(banner[ch][r])
		}
		slice = append(slice, builder.String())
	}
	return slice
}
