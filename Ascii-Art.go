package ascii

import (
	"fmt"
	"strings"
)

const (
	alpha = "ABCDEFGHIJKLMNOPQRSTUVWXYZ?"
	L     = 4
	H     = 5
)

var text1 []string = []string{
	" #  ##   ## ##  ### ###  ## # # ###  ## # # #   # # ###  #  ##   #  ##   ## ### # # # # # # # # # # ### ###",
	"# # # # #   # # #   #   #   # #  #    # # # #   ### # # # # # # # # # # #    #  # # # # # # # # # #   #   #",
	"### ##  #   # # ##  ##  # # ###  #    # ##  #   ### # # # # ##  # # ##   #   #  # # # # ###  #   #   #   ##",
	"# # # # #   # # #   #   # # # #  #  # # # # #   # # # # # # #    ## # #   #  #  # # # # ### # #  #  #      ",
	"# # ##   ## ##  ### #    ## # # ###  #  # # ### # # # #  #  #     # # # ##   #  ###  #  # # # #  #  ###  # ",
}

type ascii struct {
	x int
	y int
}

func find(c rune) int {
	for i, v := range alpha {
		if v == c {
			return i
		}
	}
	return -1
}
func Print(asciis []ascii, j int) {
	for _, v := range asciis {
		for k := v.x; k < v.y; k++ {
			fmt.Print(string(text1[j][k]))
		}
	}
	fmt.Println()
}
func PrintAscii(T string) {
	T = strings.ToUpper(T)
	var asciis []ascii
	for _, v := range T {
		j := find(v)
		var newascii ascii
		newascii.x = j * L
		newascii.y = (j + 1) * L
		asciis = append(asciis, newascii)
	}
	for i := 0; i < H; i++ {
		Print(asciis, i)
	}
}
