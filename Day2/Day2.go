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

		check := checkSafe(line)
		if check == nil {
			count++
		}

	}

	fmt.Println(count)
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
