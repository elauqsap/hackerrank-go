package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var size, count int
	var grades []int
	fmt.Scanf("%d", &size)
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() && (count < size) {
		grade, _ := strconv.Atoi(scanner.Text())
		switch {
		case grade < 38:
			fmt.Sprintf("%d < 38", grade)
			grades = append(grades, grade)
		case grade >= 38:
			fmt.Sprintf("%d >= 38", grade)
			grades = append(grades, roundGrade(grade))
		}
		count++
	}
	for _, grade := range grades {
		fmt.Println(grade)
	}
}

func roundGrade(grade int) int {
	for i := 0; i < 2; i++ {
		grade++
		if (grade % 5) == 0 {
			return grade
		}
	}
	return grade - 2
}
