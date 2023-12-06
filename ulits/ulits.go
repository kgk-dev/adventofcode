package ulits

import (
	"aoc/types"
	"bufio"
	"io"
	"log"
)

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
