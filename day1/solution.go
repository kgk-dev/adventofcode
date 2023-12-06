package day1

import (
	"bufio"
	"io"
	"log"
	"os"
	"regexp"
	"strconv"
)

func inputProducer(file *os.File) <-chan string {
	inputStream := make(chan string)
	inputReader := bufio.NewReader(file)
	go func() {
		defer close(inputStream)
		for {
			line, _, err := inputReader.ReadLine()
			if err != nil {
				if err == io.EOF {
					break
				}
				log.Fatal(err)
			}
			inputStream <- string(line)
		}
	}()
	return inputStream
}

func calc(
	inputStream <-chan string,
	regex *regexp.Regexp) <-chan int64 {
	var outputStream = make(chan int64)
	go func() {
		defer close(outputStream)
		for line := range inputStream {
			nums := regex.FindAllString(line, -1)
			digit, _ := strconv.ParseInt(nums[0]+nums[len(nums)-1], 10, 64)
			outputStream <- digit
		}
	}()
	return outputStream
}

func Solution() int64 {
	var result int64
	var regex = regexp.MustCompile(`[0-9]`)
	file, err := os.OpenFile("./day1/input", os.O_RDONLY, 0664)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	for digit := range calc(inputProducer(file), regex) {
		result += digit
	}
	return result
}
