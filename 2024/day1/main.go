package main

import (
	"bufio"
	"fmt"
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

		reportSafe, nums, dSlot := ProcessReport(input)

		if reportSafe {
			safeReports++
		}
		fmt.Printf("[%d]{%t} - nums%d\n", dSlot, reportSafe, nums)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
	fmt.Println(safeReports)
}

func ProcessReport(input []string) (bool, []int, int) {
	dampenerSlot := -1
	nums := make([]int, len(input))
	for i, v := range input {
		nums[i], _ = strconv.Atoi(v)
	}

	// do diffs
	_, dampenerSlot, diffFail := CheckDiffs(nums, dampenerSlot)
	if diffFail {
		return false, nums, dampenerSlot
	}

	return true, nums, dampenerSlot
}

func CheckDiffs(n []int, dslot int) ([]int, int, bool) {
	diffFail := false
	d := make([]int, len(n))

	for i, v := range n {
		if dslot == -1 {

		} else {
			if i == dslot+1 {
				if dslot > 0 {
					d[i] = n[i-2] - n[i]
					if d[i] > 3 || d[i] < 1 {
						diffFail = true
					}
				}
			} else {
				d[i] = n[i-1] - n[i]
				if d[i] > 3 || d[i] < 1 {
					diffFail = true
				}
			}
		}
	}

	return d, dslot, diffFail
}

func AbsInt(i1 int, i2 int) int {
	if i1 > i2 {
		return i1 - i2
	}
	return i2 - i1
}
