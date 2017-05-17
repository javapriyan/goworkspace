package main

func Pic(dx, dy int) [][]uint8 {
	x := make([][]uint8, dy)
	for i, ele := range x {
		ele = make([]uint8, dx)
		for j, _ := range ele {
			ele[j] = uint8(j)
		}

		x[i] = ele
	}
	return x
}
