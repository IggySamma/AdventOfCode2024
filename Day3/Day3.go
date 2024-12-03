package main

import (
	"fmt"
	"regexp"

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

	regEx, _ := regexp.Compile("[mul(]{1}[0-9]{1,3}[,]{1}[0-9]{1,3}[)]{1}")

	matches := regEx.FindAllStringSubmatchIndex(string(file), -1)

	fmt.Println(matches)
	fmt.Println(matches[0])
	fmt.Println(matches[0][0])
	fmt.Println(matches[0][1])
	/*
		for i := 0; i < len(matches); i++ {

		}*/
}
