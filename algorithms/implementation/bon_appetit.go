package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var n, k, b, a int
	fmt.Scanf("%d %d\n", &n, &k)
	items := make([]int, n, n)
	scanner := bufio.NewScanner(os.Stdin)
	for i := 0; i < 2; i++ {
		scanned := scanner.Scan()
		if i == 0 && scanned {
			stringToIntSlice(scanner.Text(), items)
		} else if scanned {
			b, _ = strconv.Atoi(scanner.Text())
		}
	}
	for index, val := range items {
		if index != k {
			a = a + val
		}
	}
	if (a / 2) < b {
		fmt.Println(b - (a / 2))
	} else {
		fmt.Println("Bon Appetit")
	}
}

func stringToIntSlice(s string, slice []int) {
	for index, val := range strings.Split(s, " ") {
		if i, err := strconv.Atoi(val); err == nil {
			slice[index] = i
		}
	}
}
