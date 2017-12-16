package main

import "fmt"

func main() {
	var x1, v1, x2, v2 int
	fmt.Scanf("%d %d %d %d", &x1, &v1, &x2, &v2)
	if kangarooJump(x1, v1, x2, v2) {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}

func kangarooJump(x1 int, v1 int, x2 int, v2 int) bool {
	if (x2 > x1 && v2 > v1) || (v1 == v2) {
		return false
	} else if (x2-x1)%(v1-v2) == 0 {
		return true
	}
	return false
}
