package day2

import (
	"aoc/types"
	"aoc/utlis"
	"regexp"
	"strconv"
	"strings"
)

const (
	R = 12
	G = 13
	B = 14
)

func clac[T int64](regex *regexp.Regexp) types.Calc[T] {
	var lineNo T = 0
	return func(line string) T {
		lineNo++
		for _, pair := range regex.FindAllStringSubmatch(string(line), -1) {
			splitedPair := strings.Split(pair[0], " ")
			num, _ := strconv.ParseInt(splitedPair[0], 10, 64)
			color := splitedPair[1]
			if color == "r" && num > R ||
				color == "g" && num > G ||
				color == "b" && num > B {
				return 0
			}
		}
		return lineNo
	}
}

func Solution() int64 {
	regex := regexp.MustCompile(`\d+\s+[r|g|b]`)
	return utlis.FileToResult[int64]("./day2/input",
		clac(regex),
		utlis.Add,
	)
}
