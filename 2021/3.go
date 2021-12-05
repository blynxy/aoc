package main

import (
	"fmt"
	"strconv"
)
var bitlength = 5

func main() {
	contents := FileByLine("3.txt")
	//contents := FileByLine("3_sample.txt")
	bitlength = len(contents[0])
	fmt.Println("length: ", bitlength)

	var data []int
	for _, val := range contents {
		n,err := strconv.ParseInt(val,2, 64)
		if err != nil {
			fmt.Println("DEAD")
		}
		//fmt.Printf("%05s\n", strconv.FormatInt(n,2))
		data = append(data, int(n))
	}
	fmt.Println(len(contents), " " , len(data))


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
	fmt.Printf("z/o: [%d][%d]\n",len(z),len(o))
	fmt.Printf("l/m: [%d][%d]\n",len(lcb),len(mcb))
	// lcb
	for i:=1;i<bitlength;i++ {
		z,o = SplitOnPostion(lcb, i)
		fmt.Printf("z/o: [%d][%d]\n",len(z),len(o))
		if len(z) > len(lcb)/2 {
			lcb = o
			fmt.Printf("     (o)\n")
		} else {
			lcb = z
			fmt.Printf("     (z)\n")
		}

		fmt.Printf("l: [%d]\n",len(lcb))
		if len(lcb) == 1 {
			break
		}
	}
	println("end lcb - - - - - -")
	// mcb
	for i:=1;i<bitlength;i++ {
		z,o = SplitOnPostion(mcb, i)
		fmt.Printf("z/o: [%d][%d]\n",len(z),len(o))
		if len(z) > len(mcb)/2 {
			mcb = z
			fmt.Printf("     (z)\n")
		} else {
			mcb = o
			fmt.Printf("     (o)\n")
		}

		fmt.Printf("l: [%d]\n",len(mcb))
		if len(mcb) == 1 {
			break
		}
	}

	return lcb[0],mcb[0]
}
func SplitOnPostion(d []int, pos int) ([]int, []int) {
	//fmt.Println("NEW SPLITTING ____")
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
	//fmt.Printf("mask: %012s\n", strconv.FormatInt(int64(mask), 2))
	//for _,v := range ones {
	//	fmt.Printf("o---: %012s [%d]\n", strconv.FormatInt(int64(v), 2),v)
	//}
	//fmt.Printf("mask: %012s\n", strconv.FormatInt(int64(mask), 2))
	//for _,v := range zeros {
	//	fmt.Printf("z---: %012s [%d]\n", strconv.FormatInt(int64(v), 2),v)
	//}
	return zeros, ones
}