package main

import (
	"fmt"
	"log"
	"math"
	"math/rand"
	"os"
	"strconv"
	"time"

	"github.com/schollz/progressbar/v3"
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
	bar := progressbar.Default(int64(points))

	for c := 0; c <= points; c++ {
		x = rand.Float64()
		y = rand.Float64()

		if isCircle(x, y) {
			inCircle++
		}
		PI = 4.0 * float64(inCircle) / float64(c)
		bar.Add(1)
	}
	fmt.Printf("PI: %v (Accurate: %v%%)", PI, (PI/math.Pi)*100)

}

func isCircle(x float64, y float64) bool {
	// (x - x0)^2 + (y - y0)^2 = r^2
	// assuming center is (x0, y0)
	// Circle is shifted to the right (x0 = 1, y0 = 0); r = 1
	coords := math.Sqrt((x-1)*(x-1) + y*y)
	if coords <= 1 {
		//return TRUE if point is inside circle
		return true
	} else {
		return false
	}
}
