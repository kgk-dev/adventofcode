package day1

import (
	"aoc/types"
	"aoc/ulits"
	"log"
	"os"
	"regexp"
	"strconv"
)

func calc[T int64](regex *regexp.Regexp) types.Calc[T] {
	return func(line string) T {
		nums := regex.FindAllString(line, -1)
		digit, _ := strconv.ParseInt(nums[0]+nums[len(nums)-1], 10, 64)
		return T(digit)
	}
}

func Solution() int64 {
	var result int64
	var regex = regexp.MustCompile(`[0-9]`)

	file, err := os.OpenFile("./day1/input", os.O_RDONLY, 0664)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	for r := range ulits.TransferStream(ulits.ReadLineStream(file), calc(regex)) {
		result += r 
	}

	return result
}
