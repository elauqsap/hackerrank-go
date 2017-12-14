package main

import (
	"fmt"
)

func main() {
	var alice1, alice2, alice3 int
	var bob1, bob2, bob3 int
	var aliceScore, bobScore int

	fmt.Scanf("%d %d %d\n%d %d %d", &alice1, &alice2, &alice3, &bob1, &bob2, &bob3)
	alice := []int{alice1, alice2, alice3}
	bob := []int{bob1, bob2, bob3}
	for index, score := range alice {
		switch {
		case score > bob[index]:
			aliceScore++
		case score < bob[index]:
			bobScore++
		}
	}
	fmt.Printf("%d %d", aliceScore, bobScore)
}
