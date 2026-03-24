package main

import "fmt"

func itoa(i int)string {
	str := ""
	for i > 0 {
		digit := i%10
		str = string(rune(digit) + '0') + str
		i = i/10
	}
	return str
}
func main(){
	fmt.Printf("%q\n", itoa(12345))
}