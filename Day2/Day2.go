package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	count := 0
	count2 := 0

	for scanner.Scan() {
		raw := strings.Split(scanner.Text(), " ")
		var line []int64

		for i := 0; i < len(raw); i++ {
			temp, err := strconv.ParseInt(raw[i], 10, 64)
			if err != nil {
				log.Fatal(err)
			}
			line = append(line, temp)
		}
		// Part 1
		check := checkSafe(line)
		if check == nil {
			count++
		}

		//Part 2
		check2 := checkSafe2(line)
		if check2 == nil {
			count2++
		}

	}

	fmt.Println("part 1: ", count)
	fmt.Println("part 2: ", count2)

}

func checkSafe(line []int64) error {
	if (line[1] - line[0]) > 0 {
		for i := 0; i < len(line)-1; i++ {
			if !((line[i+1]-line[i]) > 0 && (line[i+1]-line[i]) < 4) {
				return errors.New("1")
			}
		}
	} else if (line[1] - line[0]) < 0 {
		for i := 0; i < len(line)-1; i++ {
			if !((line[i+1]-line[i]) < 0 && (line[i+1]-line[i]) > -4) {
				return errors.New("1")
			}
		}
	} else {
		return errors.New("1")
	}
	return nil
}

func remove(list []int64, i int) []int64 {
	return append(list[:i], list[i+1:]...)
}

func checkSafe2(line []int64) error {
	check := checkSafe(line)
	if check == nil {
		return nil
	}
	for i := 0; i < len(line); i++ {
		var newLine []int64
		newLine = append(newLine, line...)
		newLine = remove(newLine, i)
		check2 := checkSafe(newLine)
		if check2 == nil {
			return nil
		}
	}
	return errors.New("1")
}
