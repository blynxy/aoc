package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
)

func main() {
	file, err := os.Open("data.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	safeReports := 0

	for scanner.Scan() {
		reg := regexp.MustCompile("\\s+")
		input := reg.Split(scanner.Text(), -1)

		reportSafe, nums, diffs, dSlot := ProcessReport(input)

		if reportSafe {
			safeReports++
		}
		fmt.Println(reportSafe, dSlot, nums, diffs)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
	fmt.Println(safeReports)
}

func ProcessReport(input []string) (bool, []int, []int, int) {
	dampenerSlot := -1
	nums := make([]int, len(input))
	for i, v := range input {
		nums[i], _ = strconv.Atoi(v)
	}

	diffStatus, dampenerSlot := ProcessDiffs(nums, dampenerSlot)

	pos := 0
	neg := 0
	for i, v := range diffs {
		if i == dampenerSlot {

		} else if v > 0 {
			pos++
		} else {
			neg++
		}
	}
	isIncreasing := neg > pos

	for i, v := range diffs {
		if i > 0 {
			if i == dampenerSlot {
			} else {
				if v < 0 && !isIncreasing || v > 0 && isIncreasing {
					if dampenerSlot != -1 {
						return false, nums, diffs, dampenerSlot
					} else {
						dampenerSlot = i
					}
				}

			}

		}
	}
	diffs = GetDiffs(nums, dampenerSlot)

	return true, nums, diffs, dampenerSlot
}

func ValidDiff(d int) bool {
	return Abs(d) > 0 && Abs(d) < 4
}

func ProcessDiffs(n int[], dSlot int) bool, int {
	diffs := GetDiffs(n, dSlot)
	for i,v := range diffs{
		if !ValidDiff(v) {
			if dSlot != -1 {
				return false, dSlot
			}
			if i > 2{
				dSlot = i
			} else if i == 1 {
				if ValidDiff()
			}
		}
	}

	return true, dSlot
}

func GetDiffs(n []int, dampener int) []int {
	d := make([]int, len(n))
	for i, _ := range n {
		if i == 0 {
			d[i] = 1
		} else if i == dampener {
			d[i] = 1
		} else {
			if i == dampener+1 && i > 1 {
				d[i] = n[i-2] - n[i]
			} else {
				d[i] = n[i-1] - n[i]
			}
		}
	}

	return d
}

func Abs(i int) int {
	return int(math.Abs(float64(i)))
}
func AbsInt(i1 int, i2 int) int {
	if i1 > i2 {
		return i1 - i2
	}
	return i2 - i1
}
