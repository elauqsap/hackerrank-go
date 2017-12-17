package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var n, k, count int
	fmt.Scanf("%d %d\n", &n, &k)
	pairs := make([]int, n, n)
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		stringToIntSlice(scanner.Text(), pairs)
	}

	for i := 0; i < n; i++ {
		for index, val := range pairs {
			if index != i && ((val+pairs[i])%k) == 0 && index >= i {
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
