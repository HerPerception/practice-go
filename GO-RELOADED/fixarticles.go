package main

import "strings"

func AToAn(s string) string {
	str := strings.Fields(s)
	for i := 0; i < len(str)-1; i++ {
		if str[i] == "a" && strings.ContainsAny(string(str[i+1][0]), "aeiouhAEIOH") {
			str[i] = "an"
		} else if str[i] == "A" && strings.ContainsAny(string(str[i+1][0]), "aeiouhAEIOH") {
			str[i] = "An"
		}
	}
	return strings.Join(str, " ")
}

// func FixQuotes(strs string) string {
// 	str := strings.Fields(strs)
// 	s := strings.Join(str, " ")
// 	flag := 0
// 	text := ""
// 	for i := 0; i < len(s); i++ {
// 		if i > 0 && s[i] == ' ' && s[i-1] == '\'' {
// 			continue
// 		} else if i < len(s)-1 && s[i] == ' ' && s[i+1] == '\'' {
// 			continue
// 		}
// 		if s[i] == '\'' && flag == 0 {
// 			text += " " + string(s[i])
// 			flag = 1
// 		} else if s[i] == '\'' && flag == 1 {
// 			text += string(s[i]) + " "
// 			flag = 0
// 		} else {
// 			text += string(s[i])
// 		}
// 	}
// 	return text
// }
