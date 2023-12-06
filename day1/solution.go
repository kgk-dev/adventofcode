package day1

import (
	"aoc/types"
	"aoc/ulits"
	"regexp"
	"strconv"
)

func calc(regex *regexp.Regexp) types.Calc[int64] {
	return func(line string) int64 {
		nums := regex.FindAllString(line, -1)
		digit, _ := strconv.ParseInt(nums[0]+nums[len(nums)-1], 10, 64)
		return digit
	}
}

func Solution() int64 {
	var regex = regexp.MustCompile(`[0-9]`)
	return ulits.FileToResult[int64]("./day2/input",
		calc(regex),
		ulits.Add,
	)
}
