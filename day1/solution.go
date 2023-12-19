package day1

import (
	"aoc/types"
	"aoc/utlis"
	"strings"
)

func calc() types.Calc[int64] {
	digitMap := map[string]string{
		"one":   "1",
		"two":   "2",
		"three": "3",
		"four":  "4",
		"five":  "5",
		"six":   "6",
		"seven": "7",
		"eight": "8",
		"nine":  "9",
	}
	return func(line string) int64 {
		nums := make([]string, 0)
		for i, l := range line {
			for k, v := range digitMap {
				if strings.HasPrefix(line[i:], k) {

					nums = append(nums, v)
				}
			}
			if utlis.IsDigit(l) {
				nums = append(nums, string(l))
			}
		}
		return utlis.ToInteger(nums[0] + nums[len(nums)-1])
	}
}

func Solution() int64 {
	return utlis.FileToResult[int64]("./day1/input",
		calc(),
		utlis.Add,
	)
}
