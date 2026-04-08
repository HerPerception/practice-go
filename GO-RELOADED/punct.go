package main

import "strings"

func FixPunct(s string) string {
	str := strings.Fields(s)
	strs := strings.Join(str, " ")
	text := ""
	for i := 0; i < len(strs); i++ {
		if strs[i] == ' ' && strings.ContainsAny(string(strs[i+1]), "!.,;:?") {
			continue
		} else {
			text += string(strs[i])
		}
	}
	return text
}
