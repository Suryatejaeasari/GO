package basics

import (
	"fmt"
)

// Pattern9 prints a diamond pattern of stars.
func Pattern9(n int) {
	// Top half of the diamond

	for i := 1; i <= n; i++ {

		j := n - i
		for j > 0 {
			fmt.Printf(" ")
			j--
		}
		f := i*2 - 1
		for f > 0 {
			fmt.Printf("*")
			f--
		}
		fmt.Println()
	}

	for i := n; i > 0; i-- {
		f := n - i
		for f > 0 {
			fmt.Printf(" ")
			f--
		}
		j := i*2 - 1
		for j > 0 {
			fmt.Printf("*")
			j--
		}

		fmt.Println()
	}
}
