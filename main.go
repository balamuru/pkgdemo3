package main

import (
	"fmt"
	"example.local/pkgdemo3/math"
	"example.local/pkgdemo3/math/stats"

	
)

func main() {
	fmt.Println("Hello World")
	a := math.Add(3,1)
	s := math.Subtract(3,1)
	fmt.Printf("sum: %d, diff: %d\n", a, s)

	vals := []float64{1,2,3,4,5}
	avg := stats.Avg(vals)
	fmt.Printf("avg: %f\n", avg)


}
