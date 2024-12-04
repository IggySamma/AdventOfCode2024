package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Scan()

	var matrix []string
	for scanner.Scan() {
		matrix = append(matrix, scanner.Text())
	}

	//140 width
	//139 length
	//search xmas 4 chars
	l2r := 0
	r2l := 0
	t2b := 0
	/*b2t := 0
	l2rd := 0
	r2ld := 0
	b2tl2rd := 0
	b2tr2ld := 0*/

	var tMatrix []string

	for i := 0; i < len(matrix); i++ {
		if i+3 < len(matrix) {
			for j := i; j < i+3; j++ {
				tMatrix = append(tMatrix, matrix[j])
			}
			t2b = t2b + top2bottom(tMatrix)
		}

		l2r = l2r + left2right(matrix[i])
		r2l = r2l + right2left(matrix[i])
		tMatrix = nil
	}

	fmt.Println("left to right", l2r)
	fmt.Println("right to left", r2l)

}

func top2bottom(matrix []string) int {
	found := 0

	return found
}

func left2right(matrix string) int {
	found := 0

	for i := 0; i < len(matrix)-1; i++ {
		if i+3 < len(matrix)-1 {
			if string(matrix[i:i+4]) == "XMAS" {
				found++
			}
		}
	}

	return found
}

func right2left(matrix string) int {
	found := 0

	for i := len(matrix) - 1; i >= 0; i-- {
		if i-3 > 0 {
			if string(matrix[i-4:i]) == "SAMX" {
				found++
			}
		}
	}

	return found
}
