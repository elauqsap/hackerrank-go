package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var aSet, bSet []int
	var aSetSize, bSetSize int
	fmt.Scanf("%d %d", &aSetSize, &bSetSize)
	scanner := bufio.NewScanner(os.Stdin)
	for i := 0; i < 2; i++ {
		if scanner.Scan() {
			if i == 0 {
				aSet = stringToIntSlice(scanner.Text())
			} else {
				bSet = stringToIntSlice(scanner.Text())
			}
		}
	}

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
