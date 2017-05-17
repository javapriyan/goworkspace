package main

import ("fmt"
"golang.org/x/tour/pic"

// SayHi is a function that calcs given input returns the sum
func SayHi() {
	fmt.Println("Hey, you just called me")
	pic.Show()
}

// Calc is a function that calcs given input returns the sum
func Calc(x int, y int) int {
	fmt.Println(x + y)
	return x + y
}
