package basics

import (
	"fmt"
	"strings"
)

func StringFunc() {
	s1 := "A word"
	r := strings.NewReplacer("A", "Another")
	s2 := r.Replace(s1)
	fmt.Println(s2)
	fmt.Println(strings.Contains(s2, "Another"))
	fmt.Println(strings.Index(s2, "o"))
	s3 := "\n Hello \n"
	fmt.Println(s3)
	s3 = strings.TrimSpace(s3)
	fmt.Println(s3)
	fmt.Println(strings.Split("a-b-c-d", "-"))
	fmt.Println(strings.HasPrefix("SuryaEasari", "Surya"))
	fmt.Println(strings.HasSuffix("SuryaEasari", "Easari"))
	fmt.Println("Length: ", len(s2))
	fmt.Println(strings.Replace(s2, "o", "0", -1))
	fmt.Println(strings.ToLower(s3))
	fmt.Println(strings.ToUpper(s3))
}
