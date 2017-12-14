package main

import "fmt"

func main() {
	var size int
	fmt.Scanf("%d", &size)

	for i := 1; i <= size; i++ {
		for j := 0; j < size; j++ {
			if i == size-j || j > size-i {
				fmt.Print("#")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}
