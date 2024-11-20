package patternopt

import "fmt"

func Pattern22(n int) {
	s := n*2 - 1
	for i := 0; i < s; i++ {
		for j := 0; j < s; j++ {
			mD := i
			if j < mD {
				mD = j
			}
			if s-i-1 < mD {
				mD = s - i - 1
			}
			if s-j-1 < mD {
				mD = s - j - 1
			}
			fmt.Printf("%d ", n-mD)
		}
		fmt.Println()
	}
}
