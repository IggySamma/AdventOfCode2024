package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	//"bufio"
	"log"
	"os"
	//"errors"
)

func main() {
	file, err := os.ReadFile("./input.txt")
	if err != nil {
		log.Fatal(err)
	}

	//fmt.Println(string(file))
	regEx, _ := regexp.Compile(`mul\((\d{1,3}),(\d{1,3})\)`)

	matches := regEx.FindAllStringSubmatchIndex(string(file), -1)
	/*
		fmt.Println(len(matches))
		fmt.Println(matches[0])
		fmt.Println(matches[0][0])
		fmt.Println(matches[0][1])
		fmt.Println(string(file[matches[0][0]+1 : matches[0][1]-1]))
	*/

	//fmt.Println(matches)

	count := 0

	for i := 0; i < len(matches); i++ {
		raw := string(file[matches[i][0]+4 : matches[i][1]-1])
		//fmt.Println("Arrays: ", raw)
		idx := strings.Index(raw, ",")
		left, _ := strconv.Atoi(raw[0:idx])
		right, _ := strconv.Atoi(raw[idx+1:])
		count = count + (left * right)
	}

	fmt.Println("Part 1: ", count)

	regExP2, _ := regexp.Compile(`do\(\)|don't\(\)`)

	matchesP2 := regExP2.FindAllStringSubmatchIndex(string(file), -1)

	/*fmt.Println("matches: ", matches[0][0])
	fmt.Println("2nd matches: ", matchesP2[0][0])
	fmt.Println(string(file[matchesP2[0][0]:matchesP2[0][1]]))
	fmt.Println(string(file[matchesP2[1][0]:matchesP2[1][1]]))
	fmt.Println(string(file[matchesP2[2][0]:matchesP2[2][1]]))
	fmt.Println(string(file[matchesP2[3][0]:matchesP2[3][1]]))*/

	countP2 := 0
	adder := "yes"

	for i := 0; i < len(matches); i++ {
		//fmt.Println("Adder: ", adder)
		for j := len(matchesP2) - 1; j >= 0; j-- {
			if (matchesP2[j][0] < matches[i][0]) && matchesP2[j][0] > matches[i-1][0] {
				if (string(file[matchesP2[j][0]:matchesP2[j][1]])) == "do()" {
					adder = "yes"
					break
				} else {
					adder = "no"
					break
				}
			}
		}
		if adder == "yes" {
			raw := string(file[matches[i][0]+4 : matches[i][1]-1])
			//fmt.Println("Arrays: ", raw)
			idx := strings.Index(raw, ",")
			left, _ := strconv.Atoi(raw[0:idx])
			right, _ := strconv.Atoi(raw[idx+1:])
			countP2 = countP2 + (left * right)
		}

	}

	fmt.Println("Part 2: ", countP2)
}
