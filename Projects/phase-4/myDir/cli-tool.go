package myDir

import (
	"fmt"
	"strconv"
	"strings"
)

func Calculator(s string) string {
	str := strings.Fields(s)
	var output string
	switch str[0] {
	case "add":
		first, err := strconv.Atoi(str[1])
		if err != nil {
			return "Error in first number, check input and try again"
		}
		sec, err := strconv.Atoi(str[2])
		if err != nil {
			return "Error in second number, check input and try again"
		}
		output += fmt.Sprintf("result: %d", first+sec)
		return output

	case "sub":
		first, err := strconv.Atoi(str[1])
		if err != nil {
			return "Error in first number, check input and try again"
		}
		sec, err := strconv.Atoi(str[2])
		if err != nil {
			return "Error in second number, check input and try again"
		}
		return fmt.Sprintf("result: %d", first-sec)

	case "mul":
		first, err := strconv.Atoi(str[1])
		if err != nil {
			return "Error in first number, check input and try again"
		}
		sec, err := strconv.Atoi(str[2])
		if err != nil {
			return "Error in second number, check input and try again"
		}
		return fmt.Sprintf("result: %d", first*sec)

	case "mod":
		first, err := strconv.Atoi(str[1])
		if err != nil {
			return "Error in first number, check input and try again"
		}
		sec, err := strconv.Atoi(str[2])
		if err != nil {
			return "Error in second number, check input and try again"
		}
		return fmt.Sprintf("result: %d", first%sec)

	case "pow":
		first, err := strconv.Atoi(str[1])
		if err != nil {
			return "Error in first number, check input and try again"
		}
		sec, err := strconv.Atoi(str[2])
		if err != nil {
			return "Error in second number, check input and try again"
		}
		result := 1
		for i := 0; i < sec; i++ {
			result *= first
		}
		return fmt.Sprintf("result: %d", result)

	}
	if str[0] == "last" {
		if output != "" {
			return output
		} else {
			return "No previous session. Type in a command to begin."
		}
	}

	return "Invalid command. Check input and try again."
}
