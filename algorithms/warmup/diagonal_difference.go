package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	var matrix map[int][]int
	first := true
	var size, index int
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		if first {
			size, _ = strconv.Atoi(scanner.Text())
			matrix = make(map[int][]int, size)
			first = false
		} else {
			matrix[index] = stringToIntSlice(scanner.Text())
			index++
		}
	}
	fmt.Printf("%d\n", diagonalDiff(matrix, size))
}

func diagonalDiff(matrix map[int][]int, size int) (diff int) {
	primary := 0
	secondary := 0
	for i := 0; i < size; i++ {
		primary = primary + matrix[i][i]
		secondary = secondary + matrix[i][(size-1)-i]
	}
	return int(math.Abs(float64(primary - secondary)))
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
