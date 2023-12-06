package day2

import (
	"aoc/types"
	"aoc/ulits"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

const (
	R     = 12
	G     = 13
	B     = 14
	INPUT = `Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue
Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red
Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red
Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green
`
)

func clac(regex *regexp.Regexp) types.Calc[int] {
	lineNo := 0
	return func(line string) int {
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

func Solution() int {
	file, err := os.Open("./day2/input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	regex := regexp.MustCompile(`\d+\s+[r|g|b]`)
	result := 0
	for r := range ulits.TransferStream(ulits.ReadLineStream(file), clac(regex)) {
		result += r
	}

	return result
}
