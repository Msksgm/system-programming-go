package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	fmt.Fprintf(os.Stdout, "Write with os.Stdout at %v\n", time.Now())
	fmt.Fprintf(os.Stdout, "Write with os.Stdout at %d\n", 1)
	fmt.Fprintf(os.Stdout, "Write with os.Stdout at %s\n", "string")
	fmt.Fprintf(os.Stdout, "Write with os.Stdout at %f\n", 0.1)
}
