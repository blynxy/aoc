package main

import (
	"fmt"
	"strconv"
)
var bitlength = 12

func main() {
	contents := FileByLine("3.txt")

	var data []int
	for _, val := range contents {
		n,_ := strconv.ParseInt(val,2, 64)
		fmt.Printf("%05s\n", strconv.FormatInt(n,2))
		data = append(data, int(n))
	}


	lcb, mcb := FindLcbAndMcb(data)
	fmt.Printf("lcb: %d | mcb: %d\ntotal: %d\n", lcb,mcb,lcb*mcb)
}

func FindLcbAndMcb(d []int) (int, int) {
	var lcb, mcb []int
	z,o := SplitOnPostion(d, 0)
	if len(z) > len(d)/2 {
		lcb, mcb = o,z
	} else {
		lcb, mcb = z,o
	}
	// lcb
	for i:=1;i<bitlength;i++ {
		z,o = SplitOnPostion(lcb, i)
		if len(z) > len(lcb)/2 {
			lcb = o
		} else {
			lcb = z
		}
		if len(lcb) == 1 {
			break
		}
	}
	// mcb
	for i:=1;i<bitlength;i++ {
		z,o = SplitOnPostion(mcb, i)
		if len(o) > len(mcb)/2 || len(mcb) == 2{
			mcb = o
		} else {
			mcb = z
		}
		if len(mcb) == 1 {
			break
		}
	}

	return lcb[0],mcb[0]
}
func SplitOnPostion(d []int, pos int) ([]int, []int) {
	fmt.Println("NEW SPLITTING ____")
	var mask int
	mask = 1 << (bitlength-1-pos)
	var ones,zeros []int
	for _,v := range d {
		if mask & v == mask {
			ones = append(ones, v)
		} else {
			zeros = append(zeros, v)
		}
	}
	fmt.Printf("mask: %012s\n", strconv.FormatInt(int64(mask), 2))
	for _,v := range ones {
		fmt.Printf("o---: %012s [%d]\n", strconv.FormatInt(int64(v), 2),v)
	}
	fmt.Printf("mask: %012s\n", strconv.FormatInt(int64(mask), 2))
	for _,v := range zeros {
		fmt.Printf("z---: %012s [%d]\n", strconv.FormatInt(int64(v), 2),v)
	}
	return zeros, ones
}