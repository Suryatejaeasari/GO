package basics

import "fmt"

func Pattern8(n int) {

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
