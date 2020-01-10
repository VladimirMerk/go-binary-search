package main

import (
	"fmt"
	"math"
)

func main() {
	//whatAmI := func(i interface{}) {
	//	switch t := i.(type) {
	//	case bool:
	//		fmt.Println("I'm a bool")
	//	case int:
	//		fmt.Println("I'm an int")
	//	default:
	//		fmt.Printf("Don't know type %T\n", t)
	//	}
	//}
	//whatAmI(true)
	//whatAmI(1)
	//whatAmI("hey")

	data := []int{3, 5, 7, 9, 12, 13, 15, 17, 20, 30}
	fmt.Println(binarySearch(data, 7))
}

func binarySearch(list []int, item int) int {
	var low, high, mid, guess, result int
	result = -1
	high = len(list) - 1
	for low <= high {
		mid = int(math.Floor(float64((low + high) / 2)))
		guess = list[mid]
		if guess == item {
			result = mid
			break
		} else if guess < item {
			low = mid + 1
		} else if guess > item {
			high = mid - 1
		}
	}
	return result
}
