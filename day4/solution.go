package day4

import (
	"aoc/types"
	"aoc/utlis"
	"regexp"
	"strings"
)

func calc() types.Calc[int64] {
	reg := regexp.MustCompile(`\d+`)
	//Part Two
	copies := make(map[int]int)
	lineNo := 1
	return func(line string) int64 {
		fullColum := strings.Index(line, ":")
		cards := strings.Split(line[fullColum+1:], "|")
		winCard := utlis.ToIntegers(reg.FindAllString(cards[0], -1))
		yourCard := utlis.ToIntegers(reg.FindAllString(cards[1], -1))
		matches := 1
		// Part Two
		for _, w := range winCard {
			for _, y := range yourCard {
				if w == y {
					matches++
				}
			}
		}
		if matches > 0 {
			copies[lineNo]++
			for line := copies[lineNo]; line > 0; line-- {
				for m := matches - 1; m > 0; m-- {
					copies[m+lineNo]++
				}
			}
		}
		lineNo++
		return int64(copies[lineNo-1])
		// Part One
		// matches := 0
		// for _, w := range winCard {
		// 	for _, y := range yourCard {
		// 		if w == y {
		// 			matches++
		// 		}
		// 	}
		// }
		// if matches > 0 {
		// 	return 1 << (matches - 1)
		// }
		// return 0
	}
}

func Solution() int64 {
	return utlis.FileToResult[int64](
		"./day4/input",
		calc(),
		utlis.Add,
	)
}
