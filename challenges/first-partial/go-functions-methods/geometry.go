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
type Segment struct{p []Point}

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

//Intersection validation start
func onSegment(p, q, r Point) bool{
	if (q.x <= math.Max(p.x, r.x) && q.x >= math.Min(p.x, r.x) &&
	q.y <= math.Max(p.y, r.y) && q.y >= math.Min(p.y, r.y)){
		return true
	}else{
		return false
	}
}

func orientation(p, q, r Point) int{
	val := (q.y - p.y) * (r.x - q.x) -(q.x - p.x) * (r.y - q.y)
	if val == 0 {
		return 0
	}
	
	if(val > 0){
		return 1
	}else{
		return 2
	}
}

func doIntersect(p1, q1, p2, q2 Point) bool{
	//4 orientations
	o1  := orientation(p1, q1, p2)
	o2  := orientation(p1, q1, q2)
	o3  := orientation(p2, q2, p1)
	o4  := orientation(p2, q2, q1)

	//General case
	if o1 != o2 && o3 != o4 {return true}

	//Special cases
	if (o1 == 0 && onSegment(p1, p2, q1)) {return true;}
	if (o2 == 0 && onSegment(p1, q2, q1)) {return true;}
	if (o3 == 0 && onSegment(p2, p1, q2)) {return true;}
	if (o4 == 0 && onSegment(p2, q1, q2)) {return true;}

	return false
}
//Intersection validation end

func reroll(ps []Point, n int){
	var point Point
	rand.Seed(time.Now().UnixNano())
    	min := -100
    	max := 100

		for i := 0; i < n; i++{
			newX := rand.Intn(max-min) + min
			newY := rand.Intn(max-min) + min
			point.X(float64(newX))
			point.Y(float64(newY))
			ps[i] = point
		}
}

func listSegments(ss []Segment, ps []Point, n int){
	line := make([]Point, 2)
	for i := 0; i < n-1; i++{
		line[0] = ps[i]
		line[1] = ps[i+1]
		ss[i] = Segment{line}
	}
	line[0] = ps[n-1]
	line[1] = ps[0]
	ss[n-1] = Segment{line}
}

func main() {
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
		segments := make([]Segment, sides)
		dists := make([]float64, sides)
		
		fmt.Printf("Sides: %v\n", sides)
		
		//Point Slice - start
		reroll(points, sides) //fills point slice

		for i := 0; i < sides; i++{
			fmt.Printf("Point %v: %v\n", i, points[i])
		}
		//Point Slice - end

		listSegments(segments, points, sides)

		//Intersection checking & fixing - start
		
		k := 0
		for i := 0; i < sides-1; i++{
			p1 := segments[i].p[0]
			q1 := segments[i].p[1]
			p2 := segments[i+1].p[0]
			q2 := segments[i+1].p[1]
			if(doIntersect(p1, q1, p2, q2)){
				fmt.Printf("Segment Intersection\n")
				reroll(points, sides)
				listSegments(segments, points, sides)

				for i := 0; i < sides; i++{
					fmt.Printf("Point %v: %v\n", i, points[i])
				}

				i = -1
				fmt.Printf("Attempt : %v\n", k)
				k++
				if k == 100 {
					fmt.Printf("Too many unsuccessful attempts, killing process...\n")
					break
					}				
			}else{
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
		//Intersection checking & fixing - end
		

		/*
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
		*/

	}
	
}