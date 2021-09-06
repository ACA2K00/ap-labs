//Alejandro Castro Arévalo-A01636235
package main

import (
	"math"
	"math/rand"
	"fmt"
	"os"
	"strconv"	
	"time"
)

type Point struct{ x, y float64 }

// traditional function
func Distance(p, q Point) float64 {
	return math.Hypot(q.x-p.x, q.y-p.y)
}

// same thing, but as a method of the Point type
func (p Point) Distance(q Point) float64 {
	return math.Hypot(q.x-p.x, q.y-p.y)
}

func (p *Point) X(n float64) float64{
	p.x = n
	return p.x
}

func (p *Point) Y(n float64) float64{
	p.y = n
	return p.y
}

func main() {
	var point Point
	sum := 0.0

	if len(os.Args) == 1{
		fmt.Printf("No value submitted, submit a number please... \n")
	}else{
		sides, err := strconv.Atoi(os.Args[1])
		if err != nil {
			fmt.Println(err)
			os.Exit(2)
		}
	
		points := make([]Point, sides)
		dists := make([]float64, sides)

		rand.Seed(time.Now().UnixNano())
    	min := -100
    	max := 100

		for i := 0; i < sides; i++{
			newX := rand.Intn(max-min) + min
			newY := rand.Intn(max-min) + min
			point.X(float64(newX))
			point.Y(float64(newY))
			points[i] = point
		}

		fmt.Printf("Sides: %v\n", sides)
		for i := 0; i < sides; i++{
			fmt.Printf("Point %v: %v\n", i, points[i])
		}

		for i := 0; i < sides-1; i++{
			dist := points[i].Distance(points[i+1])
			dists[i] = dist
		}
		ldist := points[sides-1].Distance(points[0])
		dists[sides-1] = ldist

		for i := 0; i < sides; i++{
			sum += dists[i]
		}
		fmt.Printf("Perimeter: %v \n", sum)
	}
}
