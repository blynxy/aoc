package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("data.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer 

	var commands []string

	// read file
	r := bufio.NewScanner(strings.NewReader(string(file)))
	r.Split(bufio.ScanRunes)

	do := string("do()")
	donot := string("don't()")
	reg := regexp.MustCompile(`(mul\(\d+,\d+\))`)

	isFirst := true
	canDo := true
	var cs string

	for r.Scan() {
		cs += r.Text()
		if len(cs) > len(do){
			if cs[:len(do)] == donot {
				canDo = true
			}
		}
		if len(cs) > len(donot){
			if cs[:len(donot)] == donot {
				
			}
		}


		fmt.Println(cs)
	}

	// process good commands
	for _, v := range commands {
		total += RunCommand(v)
	}
	fmt.Println(total)
}
func ParseCommands(s string) []string {
	var c []string
	reg := regexp.MustCompile(`mul\(\d+,\d+\)`)
	input := reg.FindAllString(s, -1)
	for i, v := range input {
		c = append(c, v)

	}
	return c
}

func RunCommand(s string) int {
	reg := regexp.MustCompile(`\d+`)
	nums := reg.FindAllString(s, -1)
	n1, _ := strconv.Atoi(nums[0])
	n2, _ := strconv.Atoi(nums[1])
	return n1 * n2
}
