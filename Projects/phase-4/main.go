package main

import (
	"bufio"
	"fmt"
	"os"

	"sentinel/myDir"
)

func main() {
	fmt.Println("SENTINEL - COMMAND & CONTROL CONSOLE")
	fmt.Println("All systems nominal. Type 'help' to begin.")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()
	// if input == "help" {
	// 	fmt.Println("calc → the calculator")
	// 	fmt.Println("base → the base converter")
	// 	fmt.Println("str → the string transformer")
	// 	fmt.Println("help → shows all commands")
	// 	fmt.Println("history → shows last 10 inputs")
	// 	fmt.Println("exit → shuts down the session")
	// } else {}
	for {
		if input == "calc" {
			fmt.Println("Choose options: add <a b>, sub <a b>, mul <a b>, div <a b>, mod <a b>, pow <a b>, last -> shows last result.")
			scanner.Scan()
			newInput := scanner.Text()
			fmt.Println(myDir.Calculator(newInput))
		}
	}
}
