package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fallen := make(map[int][]int)
	var a, b, m, n, s, t int
	var apples, oranges int
	fmt.Scanf("%d %d", &s, &t)
	fmt.Scanf("%d %d", &a, &b)
	fmt.Scanf("%d %d", &m, &n)
	scanner := bufio.NewScanner(os.Stdin)
	for i := 0; i < 2; i++ {
		if scanner.Scan() {
			fallen[i] = stringToIntSlice(scanner.Text())
		}
	}
	for fruit, distances := range fallen {
		if fruit == 0 {
			for _, d := range distances {
				if hitOrMiss((a + d), s, t) {
					apples++
				}
			}
		} else {
			for _, d := range distances {
				if hitOrMiss((b + d), s, t) {
					oranges++
				}
			}
		}
	}
	fmt.Println(apples)
	fmt.Println(oranges)
}

func hitOrMiss(d int, s int, t int) bool {
	if s <= d && d <= t {
		return true
	}
	return false
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
