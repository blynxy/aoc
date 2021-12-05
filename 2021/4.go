package main

import (
	"fmt"
	"strconv"
	"strings"
)

type BingoCard struct {
	Numbers [][]int
	Marks   [][]bool
	Won     bool
}

func (bc BingoCard) UnmarkedTotal() int {
	var sum int
	for w := 0; w < len(bc.Numbers[0]); w++ {
		for h := 0; h < len(bc.Numbers[0]); h++ {
			if !bc.Marks[w][h] {
				sum += bc.Numbers[w][h]
			}
		}
	}
	return sum
}
func (bc BingoCard) MarkCard(num int) {
	for w := 0; w < len(bc.Numbers[0]); w++ {
		for h := 0; h < len(bc.Numbers[0]); h++ {
			if bc.Numbers[w][h] == num {
				bc.Marks[w][h] = true
			}
		}
	}
}
func (bc BingoCard) CheckForWin() bool {
	// check each col
	for w := 0; w < len(bc.Numbers[0]); w++ {
		check := true
		for h := 0; h < len(bc.Numbers[0]); h++ {
			if !bc.Marks[w][h] {
				check = false
			}
		}
		if check {
			return true
		}
	}
	// check each row
	for h := 0; h < len(bc.Numbers[0]); h++ {
		check := true
		for w := 0; w < len(bc.Numbers[0]); w++ {
			if !bc.Marks[w][h] {
				check = false
			}
		}
		if check {
			return true
		}
	}
	return false
}

func main() {
	contents := FileByLine("4.txt")

	// first line is callout list
	calloutnumbers := strings.Split(contents[0], ",")

	var bingoCards []BingoCard
	var nextBingoCard BingoCard
	for i, line := range contents {
		if i < 2 {
			continue
		}
		if len(line) > 5 {
			nextRowString := strings.Fields(line)
			nextRow := []int{}
			nextMarkRow := []bool{}
			for _, v := range nextRowString {
				n, _ := strconv.ParseInt(v, 10, 64)
				nextRow = append(nextRow, int(n))
				nextMarkRow = append(nextMarkRow, false)
			}
			nextBingoCard.Numbers = append(nextBingoCard.Numbers, nextRow)
			nextBingoCard.Marks = append(nextBingoCard.Marks, nextMarkRow)
		} else {
			bingoCards = append(bingoCards, nextBingoCard)
			nextBingoCard = BingoCard{}
		}
	}
	//add last bingo card
	bingoCards = append(bingoCards, nextBingoCard)
	for i, card := range bingoCards {
		fmt.Println(card)
		bingoCards[i].Won = false
	}
	for _, callout := range calloutnumbers {
		calloutnum, _ := strconv.ParseInt(callout, 10, 64)
		fmt.Printf("~~~~ > %d < ~~~~\n", calloutnum)
		for i, card := range bingoCards {
			if !card.Won {
				card.MarkCard(int(calloutnum))
			}
			if !card.Won && card.CheckForWin() {
				fmt.Printf("WINNER! id(%d)%d\n", i, card.UnmarkedTotal()*int(calloutnum))
				bingoCards[i].Won = true
				//bingoCards = append(bingoCards[:i], bingoCards[i+1:]...)
			}
		}
	}

}
