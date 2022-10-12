package leetcode

import (
	"log"
	"math"
	"testing"
)

func TestMaxArea(t *testing.T) {
	input := [][]int{
		[]int{1, 8, 6, 20, 5, 4, 3, 20, 7},
	}
	for _, in := range input {
		out := maxArea(in)
		log.Println(in, " -> MaxArea:", out)
	}
}

func maxArea(height []int) int {
	water := 0
	j := len(height) - 1
	high := 0
	index := 1
	for i := 0; i < j; {
		log.Printf("Begin Index:%d i:%d j:%d water:%d", index, i, j, water)
		if height[i] < height[j] {
			high = height[i]
		} else {
			high = height[j]
		}
		tmpWater := int(math.Abs(float64(i-j)) * float64(high))
		if tmpWater > water {
			water = tmpWater
		}
		if height[i] < height[j] {
			i++
		} else {
			j--
		}
		log.Printf("End   Index:%d i:%d j:%d water:%d", index, i, j, water)
		index++
	}
	return water
}

func maxAreaOn2(height []int) int {
	var maxCap, tmpCap float64
	high := 0
	for i := 0; i < len(height)-1; i++ {
		for j := 0; j < len(height); j++ {
			if i == j {
				continue
			}
			if height[i] < height[j] {
				high = height[i]
			} else {
				high = height[j]
			}
			tmpCap = math.Abs(float64(i-j)) * float64(high)
			if tmpCap > maxCap {
				maxCap = tmpCap
			}
		}
	}
	return int(maxCap)
}
