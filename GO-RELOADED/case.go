package main

import (
	"strconv"
	"strings"
)

func Case(s string) string {
	str := strings.Fields(s)
	for i := 0; i < len(str); i++ {
		cmd, num := Find(str)
		if str[i] == cmd {
			if num >= i {
				num = i
			}
			for j := i - num; j < i; j++ {
				switch cmd {
				case "up":
					str[j] = strings.ToUpper(str[j])

				case "low":
					str[j] = strings.ToLower(str[j])

				case "cap":
					str[j] = strings.ToUpper(string(str[j][0])) + strings.ToLower(str[j][:1])
				}
			}
			str = append(str[:i], str[i+1:]...)
		}
	}
	return strings.Join(str, " ")
}

func Find(s []string) (string, int) {
	num := 0
	cmd := ""
	for i := 1; i < len(s); i++ {
		if strings.HasPrefix(s[i], "(") {
			cmd = s[i][1:]
			if strings.HasSuffix(s[i+1], ")") {
				temp, err := strconv.Atoi(s[i+1][:len(s[i+1])])
				if err != nil {
					return s, 0
				} else {
					num = temp
				}
			} else {
				num = 1
			}
		}
	}
	return cmd, num
}
