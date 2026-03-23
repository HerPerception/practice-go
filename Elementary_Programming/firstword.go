package main

import "fmt"
import "os"

func main() {
	var arg string = string(os.Args[1])
	// fmt.Scanln(&arg)
	// if len(arg) != 1 {
	// 	return
	// }
	// word := ""
	// for _, ch := range arg {
	// 	if ch == " " {
	// 	 break
	// 	} else {
	// 		word += string(ch)
	// 	}
	// }
	// // arg = word
	fmt.Print(arg[0])
}