package main

import (
	"bufio"
	"fmt"
	"os"
    "strings"
	"strconv"
)

func main() {
	var total int64
	first := true
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		if first {
			first = false
		} else {
            for _, sub := range strings.Split(scanner.Text(), " ") {
                i, _ := strconv.ParseInt(sub, 10, 64)
                total = total + i
            }
		}
	}
    fmt.Printf("%d\n", total)
}