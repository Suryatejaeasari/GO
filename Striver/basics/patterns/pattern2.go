package basics

import "fmt"

func Pattern2(n int) {
	for i := 1; i <= n; i++ {
		j := 1
		for j <= i {
			fmt.Printf("*")
			j++
		}
		fmt.Println()
	}
}
