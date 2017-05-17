package main

import "fmt"

/*func main() {
	fmt.Println("Welcome to go lessons")
	x, y := exampleFunc(10, 20)
	fmt.Printf("The method retuens %d and %d", x, y)
	X, Y := GetTwoNumbers_NakedFun()
	fmt.Println("The naked call params are ", X, Y)

}*/
func exampleFunc(x, y int) (int, int) {
	return y, x
}

const (
	// Create a huge number by shifting a 1 bit left 100 places.
	// In other words, the binary number that is 1 followed by 100 zeroes.
	Big = 1 << 100
	// Shift it right again 99 places, so we end up with 1<<1, or 2.
	Small = Big >> 99
)

func needInt(x int) int { return x*10 + 1 }
func needFloat(x float64) float64 {
	return x * 0.1
}

func main() {
	//pic.Show(Pic)
	//wc.Test(WordCount)
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}

}
