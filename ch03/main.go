package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"strings"
)

var csvSource = `13101,"100","1000003","トウキョウト","チヨダク","ヒトツバシ（1チョウメ）","東京都","千代田区","一ツ橋（１丁目）",1,0,1,0,0,13101,"101","1000003","トウキョウト","チヨダク","ヒトツバシ（2チョウメ）","東京都","千代田区","一ツ橋（２丁目）",1,0,1,0,0,13101,"100","1000012","トウキョウト","チヨダク","ヒトツバシ（2チョウメ）","東京都","千代田区","日比谷公園",0,0,0,0,0`

func main() {
	reader := strings.NewReader(csvSource)
	csvReader := csv.NewReader(reader)
	for {
		line, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		fmt.Println(line[2], line[6:9])
	}
}
