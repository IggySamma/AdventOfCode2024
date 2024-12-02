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

		var direction []int64

		err := checkSafe(line, direction)
		if err != nil {
			direction = nil
			break
		}

		//err := checkSafeAdded()
	}

	fmt.Println(count)
}

func checkSafeAdded(x []int64, count int) error {
	for i := 1; i < len(x); i++ {
		if x[i-1] != x[i] {
			x = nil
			return errors.New("1")
		}
	}
	count++
	return nil
}

func checkSafe(line []int64, direction []int64) error {
	for i := 0; i < len(line)-1; i++ {
		if ((line[i+1] - line[i]) < 4) && ((line[i+1] - line[i]) > -4) {
			if (line[i+1] - line[i]) > 0 {
				direction = append(direction, 1)
			} else if (line[i+1] - line[i]) < 0 {
				direction = append(direction, 0)
			}
		} else {
			return errors.New("1")
		}

	}
	return nil
}
