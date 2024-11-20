package patternopt

import (
	"fmt"
	"strings"
)

func Pattern1(n int) {
	for i := 0; i < n; i++ {
		fmt.Println(strings.Repeat("* ", n))
	}
}
