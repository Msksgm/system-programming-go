package main

import (
	"compress/gzip"
	"io"
	"os"
)

func main() {
	file, err := os.Create("multiwriter.txt")
	if err != nil {
		panic(err)
	}
	writer := gzip.NewWriter(file)
	writer.Header.Name = "text.txt"
	io.WriteString(writer, "gzip.Writer example\n")
	writer.Close()
}
