package basics

import (
	"fmt"
	"unicode/utf8"
)

func Runes() {
	a := "abcdefg"
	fmt.Println(utf8.RuneCountInString(a))
	for i, v := range a {
		fmt.Printf("%d : %#U : %c\n", i, v, v)
	}
}
