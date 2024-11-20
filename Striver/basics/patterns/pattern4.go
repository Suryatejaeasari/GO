package basics

import "fmt"

func Pattern4(n int) {
	f := 1
	for i := 1; i <= n; i++ {
		j := 1
		for j <= i {
			fmt.Printf("%d", f)
			j++

		}
		f++
		fmt.Println()
	}
}
