package main

import (
	"fmt"
	"testing"
)

func TestCalc(t *testing.T) {
	cases := []struct {
		input1, input2, expected int
	}{
		{10, 10, 20},
		{15, 15, 30},
		{16, 14, 30},
		{17, 13, 30},
		{18, 12, 30},
	}

	for _, c := range cases {
		got := Calc(c.input1, c.input2)
		if got != c.expected {
			t.Error("Error in testcase")
		}
		fmt.Printf("The sum of %d and %d is %d", c.input1, c.input2, got)
	}

}
