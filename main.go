package main

import (
	"fmt"
	"math/rand"
	"slices"
)

var max = 30

func getTwoSlices() ([]int, []int) {
	a := make([]int, 0)
	b := make([]int, 0)

	i := 0
	for i <= 15 {
		a = append(a, rand.Intn(max))
		b = append(b, rand.Intn(max))
		i+= 1
	}

	return a, b
}

func getIntersectionOfValuesInTwoClasses (slice1 []int, slice2 []int) ([]int) {
	var moreBig []int
	var minLenSlice []int
	var result []int
	bigMap := make(map[int]bool)

	if(len(slice1) > len(slice2)) {
		moreBig = slice1
		minLenSlice = slice2
	}else {
		moreBig = slice2
		minLenSlice = slice1
	}

	for _, val := range moreBig {
		bigMap[val] = true
	}

	for key := range bigMap {
		if slices.Contains(minLenSlice, key) {
			result = append(result, key)
		}
	}

	return result
}

func main() {
	a, b := getTwoSlices()
	result := getIntersectionOfValuesInTwoClasses(a, b)

	fmt.Println(result, a, b)
}