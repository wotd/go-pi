package main

import (
	"fmt"
	"log"
	"math"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func main() {

	points := 0
	var err error
	rand.Seed(time.Now().UnixNano())

	if len(os.Args) != 1 {
		points, err = strconv.Atoi(os.Args[1])
		if err != nil {
			log.Fatal(err)
		}
	} else {
		log.Fatal("Provide number points to test. Example: ./go-pi 1000")
	}

	fmt.Printf("Calculating for %v points\n", points)
	x := 0.0
	y := 0.0
	inCircle := 0
	PI := 0.0
	for c := 0; c <= points; c++ {
		x = rand.Float64()
		y = rand.Float64()

		if isCircle(x, y) {
			inCircle++
		}
		PI = 4.0 * float64(inCircle) / float64(c)
		fmt.Printf("\r[%v%%] PI: %v (Accurate: %v%%)", (c * 100 / points), PI, (PI/math.Pi)*100)
	}

}

func isCircle(x float64, y float64) bool {
	// (x - x0)^2 + (y - y0)^2 = r^2
	// assuming center is (x0, y0)
	// Circle is shifted.
	// x^2 - 2xx0 + x0^2
	coords := math.Sqrt((x-1)*(x-1) + y*y)
	if coords <= 1 {
		return true
	} else {
		return false
	}
}
