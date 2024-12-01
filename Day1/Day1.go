package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("./Input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var left []int
	var right []int

	for scanner.Scan() {
		//fmt.Println(scanner.Text())
		l, err := strconv.ParseInt(strings.Split(scanner.Text(), "   ")[0], 10, 64)
		if err != nil {
			log.Fatal(err)
		}

		r, err := strconv.ParseInt(strings.Split(scanner.Text(), "   ")[1], 10, 64)
		if err != nil {
			log.Fatal(err)
		}

		left = append(left, int(l))
		right = append(right, int(r))
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	sort.Ints(left)
	sort.Ints(right)

	part2(left, right)
	/*
		var diff []int

		for i := 0; i < 1000; i++ {
			if left[i] > right[i] {
				diff = append(diff, left[i]-right[i])
			} else {
				diff = append(diff, right[i]-left[i])
			}
		}

		sum := 0

		for i := 0; i < 1000; i++ {
			sum = sum + diff[i]
		}

		fmt.Println(sum)*/
}

func part2(left []int, right []int) {
	var sumArray []int
	var mulArray []int

	for i := 0; i < 1000; i++ {
		sumArray = append(sumArray, 0)
	}

	for i := 0; i < len(left); i++ {
		for j := 0; j < len(right); j++ {
			if int(left[i]) == int(right[j]) {
				sumArray[i] = sumArray[i] + int(1)
			}
		}
	}

	for i := 0; i < len(sumArray); i++ {
		if sumArray[i] != 0 {
			mulArray = append(mulArray, (left[i] * sumArray[i]))
		}
	}

	sum := 0

	for i := 0; i < len(mulArray); i++ {
		sum = sum + mulArray[i]
	}

	fmt.Println(sum)
}
