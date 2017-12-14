package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var max int64
	var count int
	fmt.Scanf("%d", &count)
	scanner := bufio.NewScanner(os.Stdin)
	count = 0
	if scanner.Scan() {
		for _, sub := range strings.Split(scanner.Text(), " ") {
			num, _ := strconv.ParseInt(sub, 10, 64)
			if num > max {
				max = num
				count = 1
			} else if num == max {
				count++
			}
		}
	}
	fmt.Printf("%d\n", count)
}