package main

import (
	"fmt"
	"strings"
)

// インタフェースを定義
func main() {
	var builder strings.Builder
	builder.Write([]byte("strings.Builder example\n"))
	fmt.Println(builder.String())
}
