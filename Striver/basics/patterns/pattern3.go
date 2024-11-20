package basics

import "fmt"

func Pattern3(n int) {
	for i := 1; i <= n; i++ {
		j := 1
		for j <= i {
			fmt.Printf("%d", j)
			j++
		}
		fmt.Println()
	}
}
