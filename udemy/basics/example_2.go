package basics

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func SumOfNumbers() {
	fmt.Println("Enter number: ")
	r := bufio.NewReader(os.Stdin)
	n, err := r.ReadString('\n')
	n = strings.TrimSpace(n)
	fmt.Println("Enter number: ")
	n2, err := r.ReadString('\n')
	n2 = strings.TrimSpace(n2)
	in2, err := strconv.Atoi(n2)
	in, err := strconv.Atoi(n)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d + %d = %d", in, in2, in+in2)

}
