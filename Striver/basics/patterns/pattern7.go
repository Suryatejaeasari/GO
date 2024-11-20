package basics

import "fmt"

func Pattern7(n int) {

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
}
