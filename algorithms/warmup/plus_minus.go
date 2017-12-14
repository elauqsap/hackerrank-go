package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var slice []int
	first := true
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		if first {
			first = false
		} else {
			slice = stringToIntSlice(scanner.Text())
		}
	}
	fmt.Println(plusMinus(slice))
}

func plusMinus(slice []int) string {
	var pos, neg, zero float64
	for _, num := range slice {
		switch {
		case num < 0:
			neg++
		case num > 0:
			pos++
		default:
			zero++
		}
	}
	size := float64(len(slice))
	return fmt.Sprintf("%.6f\n%.6f\n%.6f", (pos / size), (neg / size), (zero / size))
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
