package main

import (
	"fmt"
	"strconv"
)


func main() {
	contents := FileByLine("1.txt")
	var ci int
	var last1 int
	var last2 int
	var last3 int

	var curr int

	var lastsum int
	var currsum int

	for i, val := range contents {
		curr, _ = strconv.Atoi(val)
		if i == 0 {
			last1 = curr
			continue
		} else if i == 1 {
			last2 = curr
			continue
		} else if i == 2 {
			last3 = curr
			lastsum = last3+last2+last1
			continue
		}
		last1 = last2
		last2 = last3
		last3 = curr
		currsum = last3+last2+last1

		if currsum > lastsum {
			ci++
		}
		fmt.Println(val)
		lastsum = currsum
	}
	answer := ci

	fmt.Printf("Answer: %d", answer)
}