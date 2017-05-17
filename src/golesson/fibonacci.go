package main

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	x, y, z := 0, 1, 0
	return func() int {
		z = z + x
		x = y
		y = z
		return z
	}
}
