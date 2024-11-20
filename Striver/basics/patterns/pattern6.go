package basics

import "fmt"

func Pattern6(n int) {
	f := n
	for i := 1; i <= n; i++ {
		j := 1
		for j <= f {
			fmt.Printf("%d", j)
			j++
		}
		f--
		fmt.Println()
	}
}
