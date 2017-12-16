package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type lowHigh struct {
	low  int
	high int
}

func main() {
	var size int
	var scores []int
	fmt.Scanf("%d\n", &size)
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		scores = stringToIntSlice(scanner.Text())
	}
	record := lowHigh{
		low:  scores[0],
		high: scores[0],
	}
	count := lowHigh{
		low:  0,
		high: 0,
	}
	for _, score := range scores[1:size] {
		recordBreaker(score, &record, &count)
	}
	fmt.Printf("%d %d\n", count.high, count.low)
}

func stringToIntSlice(s string) []int {
	var slice []int
	for _, val := range strings.Split(s, " ") {
		if i, err := strconv.Atoi(val); err == nil {
			slice = append(slice, i)
		}
	}
	return slice
}

func recordBreaker(score int, record *lowHigh, count *lowHigh) {
	switch {
	case score < record.low:
		record.low = score
		count.low++
	case record.high < score:
		record.high = score
		count.high++
	}
}
