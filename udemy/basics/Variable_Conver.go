package basics

import (
	"fmt"
	"log"
	"reflect"
	"strconv"
)

func VarCon() {
	v1 := 1
	v2 := "hello"
	fmt.Println(reflect.TypeOf(v1), reflect.TypeOf(v2))
	v3 := "123"
	v4, err := strconv.Atoi(v3)
	v5 := strconv.Itoa(v4)
	v6 := "3.14"
	v7, err := strconv.ParseFloat(v6, 64)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(reflect.TypeOf(v3), reflect.TypeOf(v4), reflect.TypeOf(v5), reflect.TypeOf(v6), reflect.TypeOf(v7))

}
