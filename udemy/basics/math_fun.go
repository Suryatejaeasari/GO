package basics

import (
	"fmt"
	"math"
)

func Math() {
	fmt.Println("Abs(-10): ", math.Abs(-10))
	fmt.Println("Pow(4,2): ", math.Pow(4, 2))
	fmt.Println("Sqrt(4): ", math.Sqrt(4))
	fmt.Println("Cbrt(27): ", math.Cbrt(27))
	fmt.Println("Ceil(4.4): ", math.Ceil(4.4))
	fmt.Println("Floor(4.4): ", math.Floor(4.4))
	fmt.Println("Round(4.4): ", math.Round(4.4))
	fmt.Println("Log2(4): ", math.Log2(4))
	fmt.Println("Log10(100): ", math.Log10(100))
	fmt.Println("Log(10): ", math.Log(10))
	fmt.Println("Max(10,5): ", math.Max(10, 5))
	fmt.Println("Min(3,10): ", math.Min(3, 10))
	r90 := 90 * math.Pi / 180
	d90 := r90 * (180 / math.Pi)
	fmt.Println("Degree from randian: ", d90)
	fmt.Println("Sin(90): ", math.Sin(r90))
}
