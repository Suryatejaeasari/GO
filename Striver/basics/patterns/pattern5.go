package basics

import "fmt"

func Pattern5(n int) {

	for i := 0; i < n; i++ {
		j := n
		for j > i {
			fmt.Printf("*")
			j--
		}
		fmt.Println()
	}
}
