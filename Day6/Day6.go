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

	var mapp []string
	for scanner.Scan() {
		mapp = append(mapp, string(scanner.Text()))
	}

	steps := initalFind(mapp)

	fmt.Println("Part 1: ", steps)

}

func initalFind(mapp []string) []int {
	for i := 0; i < len(mapp); i++ {
		for j := 0; j < len(mapp[i]); j++ {
			if string(mapp[i][j]) == "^" {
				return move(mapp, []int{i, j})
			}
		}
	}
	return nil
}

func move(mapp []string, location []int) {

}
