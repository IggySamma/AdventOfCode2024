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
	b2t := 0
	l2rd := 0
	r2ld := 0
	b2tl2rd := 0
	b2tr2ld := 0
	xmas := 0

	for i := 0; i < len(matrix); i++ {
		if i+3 < len(matrix) {
			t2b = t2b + top2bottom(matrix, i)
			b2t = b2t + bottom2top(matrix, i)
			l2rd = l2rd + left2rightDiagonal(matrix, i)
			b2tl2rd = b2tl2rd + bottom2topLeft2right(matrix, i)
			r2ld = r2ld + right2leftDiagonal(matrix, i)
			b2tr2ld = b2tr2ld + bottom2topRight2Left(matrix, i)
		}
		if i+2 < len(matrix) {
			xmas = xmas + xMas(matrix, i)
		}

		l2r = l2r + left2right(matrix[i])
		r2l = r2l + right2left(matrix[i])
	}

	sum := l2r + r2l + t2b + b2t + l2rd + r2ld + b2tl2rd + b2tr2ld

	fmt.Println("Part 1: ", sum)

	fmt.Println("Part 2: ", xmas)

}

func xMas(matrix []string, s int) int {
	found := 0
	for i := 0; i < len(matrix[s])-2; i++ {
		if string(matrix[s+1][i+1]) == "A" {
			if string(matrix[s][i]) == "M" && string(matrix[s+2][i+2]) == "S" || string(matrix[s][i]) == "S" && string(matrix[s+2][i+2]) == "M" {
				if string(matrix[s][i+2]) == "M" && string(matrix[s+2][i]) == "S" || string(matrix[s][i+2]) == "S" && string(matrix[s+2][i]) == "M" {
					found++
				}
			}
		}
	}
	return found
}

func bottom2topRight2Left(matrix []string, s int) int {
	found := 0

	for i := 3; i < len(matrix[s]); i++ {
		if string(matrix[s][i]) == "S" && string(matrix[s+1][i-1]) == "A" && string(matrix[s+2][i-2]) == "M" && string(matrix[s+3][i-3]) == "X" {
			found++
		}
	}

	return found
}

func bottom2topLeft2right(matrix []string, s int) int {
	found := 0

	for i := 0; i < len(matrix[s])-3; i++ {
		if string(matrix[s][i]) == "S" && string(matrix[s+1][i+1]) == "A" && string(matrix[s+2][i+2]) == "M" && string(matrix[s+3][i+3]) == "X" {
			found++
		}
	}

	return found
}

func right2leftDiagonal(matrix []string, s int) int {
	found := 0

	for i := 3; i < len(matrix[s]); i++ {
		if string(matrix[s][i]) == "X" && string(matrix[s+1][i-1]) == "M" && string(matrix[s+2][i-2]) == "A" && string(matrix[s+3][i-3]) == "S" {
			found++
		}
	}

	return found
}

func left2rightDiagonal(matrix []string, s int) int {
	found := 0

	for i := 0; i < len(matrix[s])-3; i++ {
		if string(matrix[s][i]) == "X" && string(matrix[s+1][i+1]) == "M" && string(matrix[s+2][i+2]) == "A" && string(matrix[s+3][i+3]) == "S" {
			found++
		}
	}

	return found
}

func bottom2top(matrix []string, s int) int {
	found := 0

	for i := 0; i < len(matrix[s]); i++ {
		if string(matrix[s][i]) == "S" && string(matrix[s+1][i]) == "A" && string(matrix[s+2][i]) == "M" && string(matrix[s+3][i]) == "X" {
			found++
		}
	}

	return found
}

func top2bottom(matrix []string, s int) int {
	found := 0

	for i := 0; i < len(matrix[s]); i++ {
		if string(matrix[s][i]) == "X" && string(matrix[s+1][i]) == "M" && string(matrix[s+2][i]) == "A" && string(matrix[s+3][i]) == "S" {
			found++
		}
	}

	return found
}

func left2right(matrix string) int {
	found := 0

	for i := 0; i < len(matrix)-1; i++ {
		if i+3 < len(matrix) {
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
