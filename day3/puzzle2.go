package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Point struct {
	x     int
	y     int
	steps int
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// Get the various line segments from the input data
func parseInputLine(lineCoords []string) []Point {
	var line []Point
	line = append(line, Point{0, 0, 0})
	for k, v := range lineCoords {
		coord, _ := strconv.Atoi(v[1:])
		switch direction := v[0]; direction {
		case 'R':
			line = append(line, Point{line[k].x + coord, line[k].y, coord})
		case 'L':
			line = append(line, Point{line[k].x - coord, line[k].y, coord})
		case 'U':
			line = append(line, Point{line[k].x, line[k].y + coord, coord})
		case 'D':
			line = append(line, Point{line[k].x, line[k].y - coord, coord})
		}
	}

	return line
}

// Find orientation of ordered triplet
func getOrientation(p Point, q Point, r Point) int {
	var orientation = (q.y-p.y)*(r.x-q.x) - (q.x-p.x)*(r.y-q.y)

	if orientation == 0 {
		return 0
	} else if orientation > 0 {
		return 1
	} else {
		return 2
	}
}

// Does point p lie on line segment pr?
func isOnSegment(p Point, q Point, r Point) bool {
	if q.x <= max(p.x, r.x) && q.x >= min(p.x, r.x) &&
		q.y <= max(p.y, r.y) && q.y >= min(p.y, r.y) {
		return true
	}
	return false
}

// Do these two line segments intersect?
func findIntersectingSegments(p Point, q Point, p2 Point, q2 Point) []Point {
	var orientation1 = getOrientation(p, q, p2)
	var orientation2 = getOrientation(p, q, q2)
	var orientation3 = getOrientation(p2, q2, p)
	var orientation4 = getOrientation(p2, q2, q)
	intersectingSegments := []Point{p, q, p2, q2}

	// General Case
	if orientation1 != orientation2 && orientation3 != orientation4 {
		return intersectingSegments
	}

	// Special Cases
	if (orientation1 == 0 && isOnSegment(p, p2, q)) ||
		(orientation2 == 0 && isOnSegment(p, q2, q)) ||
		(orientation3 == 0 && isOnSegment(p2, p, q2)) ||
		(orientation4 == 0 && isOnSegment(p2, q, q2)) {
		return intersectingSegments
	}
	return nil
}

// Calculate the intersection point
// http://www.ambrsoft.com/MathCalc/Line/TwoLinesIntersection/TwoLinesIntersection.htm
func calculateIntersectionPoint(p Point, q Point, p2 Point, q2 Point) (Point, error) {
	denom := (q.x-p.x)*(q2.y-p2.y) - (q2.x-p2.x)*(q.y-p.y)
	if denom == 0 {
		return Point{}, errors.New("Nope")
	}
	x := ((q.x*p.y-p.x*q.y)*(q2.x-p2.x) - (q2.x*p2.y-p2.x*q2.y)*(q.x-p.x)) / denom
	y := ((q.x*p.y-p.x*q.y)*(q2.y-p2.y) - (q2.x*p2.y-p2.x*q2.y)*(q.y-p.y)) / denom
	return Point{x, y, 0}, nil
}

// Find the smallest steps from a list of intersection points
func findMinimumSteps(distances []int) int {
	min := distances[0]
	for _, v := range distances {
		if v < min {
			min = v
		}
	}

	return min
}

func calculateSteps(line1 []Point, line2 []Point, i int, j int) int {
	result := findIntersectingSegments(line1[i], line1[i+1], line2[j], line2[j+1])
	if result != nil {
		res, err := calculateIntersectionPoint(result[0], result[1], result[2], result[3])
		if err != nil {
			return 0
		}
		total := 0
		//fmt.Println("Intersection Point", res)
		//fmt.Println("Intersection coords", line1[i], line1[i+1], line2[j], line2[j+1])
		total += abs(res.y-line1[i].y) + abs(res.x-line1[i].x) + abs(res.y-line2[j].y) + abs(res.x-line2[j].x)

		for k := i; k >= 0; k-- {
			total += line1[k].steps
		}
		for k := j; k >= 0; k-- {
			total += line2[k].steps
		}
		return total
	}
	return 0
}

// Find the number of steps from the intersecting points to the source for each line
func findSteps(lines [][]Point) []int {
	line1 := lines[0]
	line2 := lines[1]

	line1Size := len(lines[0])
	line2Size := len(lines[1])
	var steps []int

	if line1Size == line2Size {
		for i := 0; i <= line1Size-2; i++ {
			for j := 0; j <= line1Size-2; j++ {
				numSteps := calculateSteps(line1, line2, i, j)
				if numSteps != 0 {
					steps = append(steps, numSteps)
				}
			}
		}
	} else if line1Size > line2Size {
		for i := 0; i <= line1Size-2; i++ {
			for j := 0; j <= line2Size-2; j++ {
				numSteps := calculateSteps(line1, line2, i, j)
				if numSteps != 0 {
					steps = append(steps, numSteps)
				}
			}
		}
	} else {
		for i := 0; i < line2Size-2; i++ {
			for j := 0; j < line1Size-2; j++ {
				numSteps := calculateSteps(line1, line2, i, j)
				if numSteps != 0 {
					steps = append(steps, numSteps)
				}
			}
		}
	}
	return steps
}

func main() {
	file, err := os.Open("input2")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var lines [][]Point

	for scanner.Scan() {
		lines = append(lines, parseInputLine(strings.Split(scanner.Text(), ",")))
	}

	distances := findSteps(lines)
	fmt.Println(findMinimumSteps(distances))

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

}
