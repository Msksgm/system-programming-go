package main

import (
	"os"
)

// インタフェースを定義
func main() {
	os.Stdout.Write([]byte("os.Stdout example\n"))
}
