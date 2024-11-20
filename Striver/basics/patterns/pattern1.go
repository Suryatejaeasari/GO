package basics

import "fmt"

func Patter1(n int) {
	for i := n; i > 0; i-- {
		for j := n; j > 0; j-- {
			fmt.Printf("* ")
		}
		fmt.Println()

	}
}
