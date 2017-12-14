package main

import (
	"fmt"
	"sort"
)

func main() {
	var slice []int
	var n1, n2, n3, n4, n5 int
	fmt.Scanf("%d %d %d %d %d", &n1, &n2, &n3, &n4, &n5)
	nums := map[int]int{0: n1, 1: n2, 2: n3, 3: n4, 4: n5}
	for i := 0; i < 5; i++ {
		sum := 0
		for key, val := range nums {
			if key != i {
				sum = sum + val
			}
		}
		slice = append(slice, sum)
	}
	sort.Ints(slice)
	fmt.Printf("%d %d\n", slice[0], slice[4])
}
