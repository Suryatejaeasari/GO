package patternopt

import (
	"fmt"
)

//	func Pattern11(n int) {
//		for i := 1; i <= n; i++ {
//			s := i % 2
//			var sb strings.Builder
//			for j := 0; j < i; j++ {
//				sb.WriteByte(byte('0' + (s+j)%2))
//			}
//			fmt.Println(sb.String())
//		}
//	}
func Pattern11(n int) {
	for i := 1; i <= n; i++ {
		s := i % 2
		for j := 1; j <= i; j++ {
			fmt.Printf("%d", (s+j)%2)
		}
		fmt.Println()
	}
}
