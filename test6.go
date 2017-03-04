package main

import (
	"fmt"
	"time"
	"math/rand"
)


// this is leetcode medium level question 11.
// Given n non-negative integers a1, a2, ..., an, where each represents a point at coordinate (i, ai).
// n vertical lines are drawn such that the two endpoints of line i is at (i, ai) and (i, 0). 
// Find two lines, which together with x-axis forms a container,
// such that the container contains the most water.

Note: You may not slant the container and n is at least 2.
func main() {
	s := time.Now()
	a := []int{}
	for i:=0;i<10000;i++{
		r:=rand.Intn(100)
		a = append(a, r)
	}
	//fmt.Println(a)
	maxArea(a)
	e:=time.Since(s)
	fmt.Printf("%s", e)
}

func maxArea(height []int) int {
	//for all possible combinations, bruteforce
	maxArea := 0
	maxI := 0
	maxJ := 0
	for i := 0; i < len(height); i++ {
		for j := i + 1; j < len(height); j++ {
			next := getArea(height[i], height[j], j-i)
			if next >= maxArea {
				maxArea = next
				maxI = i
				maxJ = j
			}
		}
	}
	fmt.Println("maxArea:", maxArea, maxI, maxJ)

	//or just zero in from two sides
	return maxArea
}

func getArea(h1 int, h2 int, w int) (area int) {
	if h1 > h2 {
		return w * h2
	} else {
		return w * h1
	}
}
