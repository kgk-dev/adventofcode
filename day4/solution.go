package day4

import (
	"aoc/types"
	"aoc/utlis"
	"fmt"
	"strings"
)

func calc() types.Calc[int64] {
	return func(line string) int64 {
		fullColum := strings.Index(line, ":")
		cards := strings.Split(line[fullColum+1:], "|")
		winCard := cards[0]
		yourCard := cards[1]
		fmt.Println(winCard, yourCard)
		return 0
	}
}

func Solution() int64 {
	return utlis.FileToResult[int64](
		"./day4/input1",
		calc(),
		utlis.Add,
	)
}
