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
	x int
	y int
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
	line = append(line, Point{0, 0})
	count := 0
	for _, v := range lineCoords {
		coord, _ := strconv.Atoi(v[1:])
		switch direction := v[0]; direction {
		case 'R':
			line = append(line, Point{line[count].x + coord, line[count].y})
		case 'L':
			line = append(line, Point{line[count].x - coord, line[count].y})
		case 'U':
			line = append(line, Point{line[count].x, line[count].y + coord})
		case 'D':
			line = append(line, Point{line[count].x, line[count].y - coord})
		}
		count++
	}

	return line
}

//Find the manhattan distance between two points
func getManhattanDistance(x Point, y Point) int {
	return abs(y.x-x.x) + abs(y.y-x.y)
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
	return Point{x, y}, nil
}

// Find the smallest distance from a list of Manhattan distances
func findMinimumDistance(distances []int) int {
	min := distances[0]
	for _, v := range distances {
		if v < min {
			min = v
		}
	}

	return min
}

func calculateDistances(line1 []Point, line2 []Point, i int, j int) int {
	result := findIntersectingSegments(line1[i], line1[i+1], line2[j], line2[j+1])
	if result != nil {
		res, err := calculateIntersectionPoint(result[0], result[1], result[2], result[3])
		if err != nil {
			fmt.Println("Whoops", err)
			return 0
		}
		return getManhattanDistance(Point{0, 0}, res)
	}
	return 0
}

// Find the Manhattan distance of the intersecting points from the source
func findDistances(lines [][]Point) []int {
	line1 := lines[0]
	line2 := lines[1]

	line1Size := len(lines[0])
	line2Size := len(lines[1])

	var distances []int

	if line1Size == line2Size {
		for i := 0; i <= line1Size-2; i++ {
			for j := 0; j <= line1Size-2; j++ {
				distance := calculateDistances(line1, line2, i, j)
				if distance != 0 {
					distances = append(distances, distance)
				}
			}
		}
	} else if line1Size > line2Size {
		for i := 0; i <= line1Size-2; i++ {
			for j := 0; j <= line2Size-2; j++ {
				distance := calculateDistances(line1, line2, i, j)
				if distance != 0 {
					distances = append(distances, distance)
				}
			}
		}
	} else {
		for i := 0; i < line2Size-2; i++ {
			for j := 0; j < line1Size-2; j++ {
				distance := calculateDistances(line1, line2, i, j)
				if distance != 0 {
					distances = append(distances, distance)
				}
			}
		}
	}

	return distances
}

func main() {
	file, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var lines [][]Point

	for scanner.Scan() {
		lines = append(lines, parseInputLine(strings.Split(scanner.Text(), ",")))
	}

	distances := findDistances(lines)
	fmt.Println(findMinimumDistance(distances))

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

}
