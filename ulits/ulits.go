package ulits

import (
	"bufio"
	"io"
	"log"
)

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
