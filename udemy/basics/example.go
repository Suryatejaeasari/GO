package basics

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func EduSug() {
	fmt.Printf("What is your name: ")
	r := bufio.NewReader(os.Stdin)
	n, err := r.ReadString('\n')
	n = strings.TrimSpace(n)
	fmt.Printf("Hello, %s!! \nWhat's your age: ", n)
	age, err := r.ReadString('\n')
	age = strings.TrimSpace(age)
	if err != nil {
		log.Fatal(err)
	}
	ia, err := strconv.Atoi(age)
	if err != nil {
		log.Fatal(err)
	}
	if ia < 5 {
		fmt.Println("Too young to join")
	} else if ia == 5 {
		fmt.Println("Go to kindergarden")
	} else if (ia > 5) && (ia < 17) {
		i := ia - 5
		fmt.Printf("Go to %d Garde \n", i)

	} else {
		fmt.Println("Go to College")
	}

}
