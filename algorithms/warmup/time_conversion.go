package main

import (
	"fmt"
	"time"
)

func main() {
	var input string
	fmt.Scanf("%s", &input)
	parsed, _ := time.Parse("15:04:05PM", input)
	fmt.Printf("%02d:%02d:%02d\n", parsed.Hour(), parsed.Minute(), parsed.Second())
}
