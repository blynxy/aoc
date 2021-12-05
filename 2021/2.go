package main

import (
	"fmt"
	"strconv"
	"strings"
)


func main() {
	contents := FileByLine("2.txt")

	// Horizontal, Depth
	pos := [2]int{0,0}
	var aim int
	for _, val := range contents {
		directions := strings.Fields(val)
		v,_ := strconv.Atoi(directions[1])
		fmt.Println(directions[1])
		switch directions[0] {
		case "forward":
			pos[0] += v
			pos[1] += aim * v
			break
		case "down":
			aim += v
			//pos[1] += v
			break
		case "up":
			aim -= v
			//pos[1] -= v
			break
		default:
			pos[0] -= v
		}
	}
	fmt.Printf("W: %d H: %d\n",pos[0],pos[1])
	fmt.Println(pos[1]*pos[0])

	//fmt.Printf("Answer: %d", answer)
}