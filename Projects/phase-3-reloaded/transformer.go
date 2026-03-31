package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Upper(s string) string {
	str := strings.Fields(s)

	return strings.ToUpper(strings.Join(str[1:], " "))
}

func Lower(s string) string {
	str := strings.Fields(s)

	return strings.ToLower(strings.Join(str[1:], " "))
}

func Cap(s string) string {
	str := strings.Fields(s)

	for i := range str {
		str[i] = (strings.ToUpper(string(str[i][0])) + strings.ToLower(str[i][1:]))
	}

	return strings.Join(str[1:], " ")
}

func Title(s string) string {
	str := strings.Fields(s)

	for i := range str {
		str[i] = strings.ToLower(str[i])
	}
	for i := range str {
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
	}

	return strings.Join(str[1:], " ")
}

func Snake(s string) string {
	str := strings.Fields(s)

	for i := range str {
		str[i] = strings.ToLower(str[i])
	}
	word := strings.Join(str[1:], "_")
	text := ""
	for _, ch := range word {
		if ch == '_' {
			text += string(ch)
		} else if ch >= 'a' && ch <= 'z' || ch >= 'A' && ch <= 'Z' {
			text += string(ch)
		}
	}
	return text
}

func Reverse(s string) string {
	str := strings.Fields(s)

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

var history []string
var maxHistory = 5
func History(s string) string {
	if len(history) >= maxHistory {
		history = history[1:]
	}
	history = append(history, s)
	result := ""
	for i := 0; i < len(history); i++ {
		result += fmt.Sprintf("%d: %s\n", i+1, history[i])
	}
	return result
}

func main() {
	fmt.Println("Welcome To Sentinel String Transformer.")
	for {
		fmt.Println("Enter Prompt below: ")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		input := scanner.Text()

		if input == "exit" {
			fmt.Println("Shutting down String Transformer. Goodbye...")
			return
		} else if input == "history" {
			fmt.Println(History())
		} else {
			if len(input) == 0 {
				continue
			}
			input = strings.ToLower(input)
			newInput := strings.Fields(input)
			if len(newInput) < 2 {
				fmt.Println("No text provided. Usage: <command> <text>")
			}
			commandHistory := input + "=>" + output
			switch newInput[0] {
			case "upper":
				fmt.Println(Upper(input))
			case "lower":
				fmt.Println(Lower(input))

			case "cap":
				fmt.Println(Cap(input))

			case "title":
				fmt.Println(Title(input))
			case "snake":
				fmt.Println(Snake(input))

			case "reverse":
				fmt.Println(Reverse(input))

			default:
				fmt.Printf("Unknown command: %s, Valid Commands: upper, lower, cap, title, snake, reverse, exit\n", newInput[0])
			}
		}

	}
}
