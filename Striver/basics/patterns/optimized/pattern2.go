package patternopt

import (
	"fmt"
	"strings"
)

func Pattern2(n int) {
	for i := 1; i <= n; i++ {
		fmt.Println(strings.Repeat("*", i))
	}
}
