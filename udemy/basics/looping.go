package basics

import "fmt"

func Loops() {
	fmt.Println("for loop in general")
	for x := 1; x <= 5; x++ {
		fmt.Println(x)
	}
	fmt.Println("for loop as while loop")
	f := 0
	for f < 5 {
		fmt.Println(f)
		f++
	}
	fmt.Println("printing list")
	a := []int{1, 2, 3}
	for _, n := range a {
		fmt.Println(n)
	}
	fmt.Println("for loop as do while")
	v := 5
	for true {
		if v == 1 {
			fmt.Println("Bye i'm exiting!!!")
			break
		}
		fmt.Println(v)
		v--
	}

}
