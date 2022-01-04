package main

import (
	"bytes"
	"fmt"
)

// インタフェースを定義
func main() {
	var buffer bytes.Buffer
	buffer.Write([]byte("bytes.Buffer example\n"))
	fmt.Println(buffer.String())
}
