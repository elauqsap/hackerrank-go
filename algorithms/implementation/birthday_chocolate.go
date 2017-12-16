package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var size, count int
	fmt.Scanf("%d\n", &size)
	chocBar := make([]int, size, size)
	birthday := make([]int, 2, 2)
	scanner := bufio.NewScanner(os.Stdin)
	for i := 0; i < 2; i++ {
		if scanner.Scan() {
			if i == 0 {
				stringToIntSlice(scanner.Text(), chocBar)
			} else {
				stringToIntSlice(scanner.Text(), birthday)
			}
		}
	}
	for index, _ := range chocBar {
		if birthday[1]+index <= size {
			if check(chocBar[index:birthday[1]+index], birthday[0]) {
				count++
			}
		}
	}
	fmt.Println(count)
}

func stringToIntSlice(s string, slice []int) {
	for index, val := range strings.Split(s, " ") {
		if i, err := strconv.Atoi(val); err == nil {
			slice[index] = i
		}
	}
}

func check(chocoBar []int, day int) bool {
	var sum int
	for _, value := range chocoBar {
		sum = sum + value
	}
	if sum == day {
		return true
	}
	return false
}
