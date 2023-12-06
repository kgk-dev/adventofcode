package ulits

import (
	"aoc/types"
	"bufio"
	"io"
	"log"
	"os"
)

func Add(a, b int64) int64 {
	return a + b
}

func TransferStream[T any](inputStream <-chan string, fn types.Calc[T]) <-chan T {
	outputStream := make(chan T)
	go func() {
		defer close(outputStream)
		for line := range inputStream {
			outputStream <- fn(line)
		}
	}()
	return outputStream
}

func FileToResult[T any](filename string,
	calcFn types.Calc[T],
	resultFn types.Result[T]) T {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	var result T
	for r := range TransferStream[T](ReadLineStream(file), calcFn) {
		result = resultFn(result, r)
	}
	return result
}

func ReadLineStream(reader io.Reader) <-chan string {
	readerStream := make(chan string)
	bReader := bufio.NewReader(reader)
	go func() {
		defer close(readerStream)
		for {
			line, _, err := bReader.ReadLine()
			if err != nil {
				if err == io.EOF {
					break
				}
				log.Fatal(err)
			}
			readerStream <- string(line)
		}
	}()
	return readerStream
}
