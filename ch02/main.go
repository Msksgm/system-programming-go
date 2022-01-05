package main

import (
	"encoding/csv"
	"os"
)

func main() {
	file, err := os.Create("test.csv")
	if err != nil {
		panic(err)
	}
	writer := csv.NewWriter(file)
	record := []string{"hoge"}
	if err := writer.Write(record); err != nil {
		panic(err)
	}
	writer.Flush()
}
