package main

import "fmt"

func QuadA(x, y int) {
	for i := 0; i <= y-1; i++ {
		for j := 0; j <= x-1; j++ {
			if j != 0 && j != x-1 && (i == 0 || i == y-1) {
				fmt.Print("-")
			} else if i == 1 && (j == 0 || j == x-1) {
				fmt.Print("|")
			} else if i == 1 && (j != 0 || j != x-1) {
				fmt.Print(" ")
			} else if (j == 0 || j == 4 && i == 0) || (j == 0 || j == x-1 && i == y-1) {
				fmt.Print("o")
			}
		}
		fmt.Println()
	}
}

func main() {
	QuadA(5, 5)
}
