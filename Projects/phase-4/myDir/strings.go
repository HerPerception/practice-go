package myDir

import (
	"fmt"
	"strings"
)

func StringsTran(s string) string {
	str := strings.Fields(s)

	switch str[0] {
	case "upper":
		for i := range str {
			str[i] = strings.ToUpper(str[i])
		}

		return strings.Join(str[1:], " ")

	case "low":
		for i := range str {
			str[i] = strings.ToLower(str[i])
		}
		return strings.Join(str[1:], " ")

	case "cap":
		for i := 0; i < len(str); i++ {
			str[i] = (strings.ToUpper(string(str[i][0])) + strings.ToLower(str[i][1:]))
		}
		return strings.Join(str[1:], " ")

	case "title":
		for i := range str {
			str[i] = strings.ToLower(str[i])
		}
		for i := 0; i < len(str); i++ {
			if i > 1 && (str[i] == "a" || str[i] == "an" || str[i] == "the" || str[i] == "and") {
				continue
			} else if i > 1 && (str[i] == "but" || str[i] == "or" || str[i] == "for" || str[i] == "nor") {
				continue
			} else if i > 1 && (str[i] == "on" || str[i] == "at" || str[i] == "to" || str[i] == "by") {
				continue
			} else if i > 1 && (str[i] == "in" || str[i] == "of" || str[i] == "up" || str[i] == "as" || str[i] == "is" || str[i] == "it") {
				continue
			} else {
				str[i] = (strings.ToUpper(string(str[i][0])) + strings.ToLower(str[i][1:]))
			}
			return strings.Join(str[1:], " ")
		}
	case "snake":
		strs := strings.ToLower(strings.Join(str, "_"))

		word := ""
		for i := range strs {
			if strs[i] == '_' {
				word += string(strs[i])
			} else if strs[i] >= 'a' && strs[i] <= 'z' || strs[i] >= 'A' && strs[i] <= 'Z' {
				word += string(strs[i])
			}
		}
		return word

	case "reverse":
		text := ""
		for i := 0; i < len(str); i++ {
			index := str[i]
			for in := len(index) - 1; in >= 0; in-- {
				text += string(index[in])
			}
			str[i] = text
			text = ""
		}
		return strings.Join(str[1:], " ")
	}
	return fmt.Sprintf("%s is not a string transformer command. Check input and try agin.", str[0])
}
