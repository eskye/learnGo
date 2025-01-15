package main

import (
	"fmt"
	"sort"
)

func slicefunc() {
	var colors = []string{"Red", "Green", "Blue"}
	fmt.Println(colors)

	colors = append(colors, "Purple")
	fmt.Println(colors)

	colors = append(colors[1:len(colors)])
	fmt.Println(colors)

	colors = append(colors[:len(colors)-1])
	fmt.Println(colors)

	numbers := make([]int, 5, 5)
	numbers[0] = 134
	numbers[1] = 132
	numbers[2] = 114
	numbers[3] = 104
	numbers[4] = 100
	fmt.Println(numbers)

	numbers = append(numbers, 234)
	fmt.Println(numbers)

	sort.Ints(numbers)
	fmt.Println(numbers)

}
