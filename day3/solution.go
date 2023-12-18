package day3

import (
	"aoc/types"
	"aoc/ulits"
	"regexp"
)

var (
	numRegexp, _ = regexp.Compile(`\d+`)
	symRegexp, _ = regexp.Compile(`[^. | ^\d]`)
)

type Result struct {
	total int64
	prev  []byte
	cur   []byte
	next  []byte
}

func check(numberIndex []int, symbolIndexs [][]int) bool {
	var result bool
	for _, symbolIndex := range symbolIndexs {
		if numberIndex[0]-1 <= symbolIndex[0] && symbolIndex[0] <= numberIndex[1] {
			result = true
			break
		}
	}
	return result
}

func valid(prev, cur, next []byte) int64 {
	var result int64
	numbers := numRegexp.FindAllString(string(cur), -1)
	if len(numbers) > 0 {
		prevSymbolIndex := symRegexp.FindAllSubmatchIndex(prev, -1)
		curSymbolIndex := symRegexp.FindAllSubmatchIndex(cur, -1)
		nextSymbolIndex := symRegexp.FindAllSubmatchIndex(next, -1)
		curNumberIndex := numRegexp.FindAllSubmatchIndex(cur, -1)
		for i, numberIndex := range curNumberIndex {
			number := ulits.ToInteger(numbers[i])
			if check(numberIndex, prevSymbolIndex) {
				result += number
			} else if check(numberIndex, curSymbolIndex) {
				result += number
			} else if check(numberIndex, nextSymbolIndex) {
				result += number
			}
		}
	}
	return result
}

func calc(prev, cur, next []byte) types.Calc[Result] {
	return func(s string) Result {
		line := []byte(s)
		if len(cur) == 0 {
			cur = line
		} else if len(next) == 0 {
			next = line
		} else {
			prev, cur, next = cur, next, line
		}
		return Result{
			total: valid(prev, cur, next),
			prev:  prev,
			cur:   cur,
			next:  next,
		}
	}
}

func AddTotal(r1, r2 Result) Result {
	r2.total += r1.total
	return r2
}

func Solution() int64 {
	var prev, cur, next []byte
	result := ulits.FileToResult[Result]("./day3/input",
		calc(prev, cur, next),
		AddTotal,
	)

	num := valid(result.cur, result.next, []byte{})
	return num + result.total
}
