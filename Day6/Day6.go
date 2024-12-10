package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var mapp [][]string
	for scanner.Scan() {
		mapp = append(mapp, strings.Split(string(scanner.Text()), ""))
	}

	/*fmt.Println(mapp)
	fmt.Println(mapp[0])
	fmt.Println(mapp[0][0])*/

	location := initialLocation(mapp)

	fmt.Println("Start Location: ", location)

	move(mapp, location)

	check := checkX(mapp)

	//fmt.Println(mapp)

	fmt.Println("Part 1: ", check)

	path := getSteps(mapp)

	//fmt.Println("path: ", path)

	barriers := getBarriers(mapp)

	setLoops(mapp, path, barriers)

	//fmt.Println("loops: ", setLoops)

	countLoops := checkO(mapp)
	fmt.Println(mapp)
	fmt.Println("Part 2:", countLoops)

	//up, down, left, right, Path := getSteps(mapp)

	//fmt.Println("Barriers: ", barriers)
	//fmt.Println("Steps Up: ", up, " Steps Right: ", right, " Steps Down: ", down, " Steps Left: ", left)

}

func setLoops(mapp [][]string, path [][]int, barriers [][]int) {
	//Check if another barrier is on the same patch +-1 of path
	//If another line at same row/column of path would cause a redirect
	//check all 4 corners for these conditions
	//top left 0
	for i := 0; i < len(mapp); i++ {
		for j := 0; j < len(mapp[i]); j++ {
			/*
				Example:
						1,4 <- to find 					-> Barriers 0,4
						1,8 <- Array Exists				-> Barriers 1,9
						5,8 <- match left				-> Barriers 6,8
						5,4 <- return right to top find	-> Barriers 5,3
			*/
			// set to path but check for change in direction instead then do this search
			//if arrayExists([]int{i, j}, path) && multiArrayExists(j, i, path) {
			/*if arrayExists([]int{i, j}, path) {
			//if mapp[j-1][i] == "#" || mapp[j+1][i] == "#" || mapp[j][i-1] == "#" || mapp[j][i+1] == "#" {
			//if !(arrayExists([]int{j - 1, i}, barriers) || arrayExists([]int{j + 1, i}, barriers) || arrayExists([]int{j, i - 1}, barriers) || arrayExists([]int{j, i + 1}, barriers)) {
			//switch below to barriers
			/*if arrayExists([]int{j - 1, i}, barriers) {
				mapp[j-1][i] = "O"
			}
			if arrayExists([]int{j + 1, i}, barriers) {
				mapp[j+1][i] = "O"
			}
			if arrayExists([]int{j, i - 1}, barriers) {
				mapp[j][i-1] = "O"
			}
			if arrayExists([]int{j, i + 1}, barriers) {
				mapp[j][i+1] = "O"
			}*/
			/*	temp := multiArrayExists(j, i, path)
				if temp != nil {
					fmt.Println(temp)
				}
			}*/
			if mapp[i][j] == "X" {
				//fmt.Println("passed 1: ")
				if i-1 > 0 || j-1 > 0 || i+1 < len(mapp) || j+1 < len(mapp[i]) {
					//fmt.Println("passed 2: ")
					if mapp[i][j+1] == "X" && mapp[i-1][j] == "X" && mapp[i][j-1] == "X" && mapp[i+1][j] == "X" {
						//fmt.Println("passed 3: ")
						for k := 0; k < len(path); k++ {
							for l := 0; l < len(path); l++ {
								if arrayExists([]int{i, k}, path) && arrayExists([]int{l, k}, path) && arrayExists([]int{l, j}, path) ||
									arrayExists([]int{(i - 1), k}, barriers) && arrayExists([]int{(l + 1), (k)}, barriers) && arrayExists([]int{(l), (j - 1)}, barriers) {
									for m := 0; m < len(barriers); m++ {
										if !(barriers[m][0] == i && barriers[m][1] > j && barriers[m][1] < k ||
											barriers[m][0] == l && barriers[m][1] > j && barriers[m][1] < k ||
											barriers[m][1] == j && barriers[m][0] > i && barriers[m][0] < l ||
											barriers[m][1] == k && barriers[m][0] > i && barriers[m][0] < l) {
											mapp[i][j] = "O"

										} else {
											mapp[i][j] = "X"
											break
										}
									}
								}
								/*if arrayExists([]int{(i - 1), k}, barriers) && arrayExists([]int{(l + 1), (k)}, barriers) && arrayExists([]int{(l), (j - 1)}, barriers) {
								for m := 0; m < len(barriers); m++ {
									/*if !(barriers[m][0] == i && barriers[m][1] > j && barriers[m][1] < k ||
										barriers[m][0] == l && barriers[m][1] > j && barriers[m][1] < k ||
										barriers[m][1] == j && barriers[m][0] > i && barriers[m][0] < l ||
										barriers[m][1] == k && barriers[m][0] > i && barriers[m][0] < l) {
										mapp[i][j] = "O"

									} else {
										mapp[i][j] = "X"
										break
									}*/
								/*	fmt.Println("y")
										mapp[i][j] = "O"
									}
								}*/
							}
						}
					}
				}
			}

			/*if mapp[i][j] == "X" {
			if i-1 > 0 || j-1 > 0 || i+1 < len(mapp) || j+1 < len(mapp[i]) {
				for k := 0; k < len(barriers); k++ {
					for l := 0; l < len(barriers); l++ {
						if arrayExists([]int{(i - 1), k}, barriers) && arrayExists([]int{(l + 1), (k)}, barriers) && arrayExists([]int{(l), (j - 1)}, barriers) {
							for m := 0; m < len(barriers); m++ {
								/*if !(barriers[m][0] == i && barriers[m][1] > j && barriers[m][1] < k ||
									barriers[m][0] == l && barriers[m][1] > j && barriers[m][1] < k ||
									barriers[m][1] == j && barriers[m][0] > i && barriers[m][0] < l ||
									barriers[m][1] == k && barriers[m][0] > i && barriers[m][0] < l) {
									mapp[i][j] = "O"

								} else {
									mapp[i][j] = "X"
									break
								}*/
			/*fmt.Println("y")
									mapp[i][j] = "O"
								}
							}
						}
					}
				}
			}*/

		}
	}
}

func multiArrayExists(check int, check2 int, path [][]int) []int {
	for i := 0; i < len(path); i++ {
		for j := 0; j < len(path); j++ {
			if path[i][1] == check && path[i][0] == path[j][0] && arrayExists([]int{check, path[j][1]}, path) {
				return []int{check, path[j][1]}
			}
		}
	}
	return nil
}

/*
	func multiArrayExists(check int, check2 int, path [][]int) bool {
		for i := 0; i < len(path); i++ {
			for j := 0; j < len(path); j++ {
				if path[i][1] == check && path[i][0] == path[j][0] && arrayExists([]int{check, path[j][1]}, path) {
					return true
				}
			}
		}
		return false
	}
*/
func arrayExists(check []int, mapp [][]int) bool {
	for i := 0; i < len(mapp); i++ {
		if check[0] == mapp[i][0] && check[1] == mapp[i][1] {
			return true
		}
	}
	return false
}

func getBarriers(mapp [][]string) [][]int {
	var barriers [][]int
	for i := 0; i < len(mapp); i++ {
		for j := 0; j < len(mapp[i]); j++ {
			if mapp[i][j] == "#" {
				barriers = append(barriers, []int{i, j})
			}
		}
	}
	return barriers
}

func getSteps(mapp [][]string) [][]int {
	/*var up [][]int
	var down [][]int
	var left [][]int
	var right [][]int*/
	var path [][]int

	for i := 0; i < len(mapp); i++ {
		for j := 0; j < len(mapp[i]); j++ {
			if mapp[i][j] == "^" {
				//up = append(up, []int{i, j})
				path = append(path, []int{i, j})
			}
			if mapp[i][j] == "V" {
				//down = append(down, []int{i, j})
				path = append(path, []int{i, j})
			}
			if mapp[i][j] == "<" {
				//left = append(left, []int{i, j})
				path = append(path, []int{i, j})
			}
			if mapp[i][j] == ">" {
				//right = append(right, []int{i, j})
				path = append(path, []int{i, j})
			}
		}
	}

	return path
}

func checkX(mapp [][]string) int {
	counter := 0
	for i := 0; i < len(mapp); i++ {
		for j := 0; j < len(mapp[i]); j++ {
			if mapp[i][j] == "X" {
				counter++
			}
		}
	}
	return counter
}

func checkO(mapp [][]string) int {
	counter := 0
	for i := 0; i < len(mapp); i++ {
		for j := 0; j < len(mapp[i]); j++ {
			if mapp[i][j] == "O" {
				counter++
			}
		}
	}
	return counter
}

func initialLocation(mapp [][]string) []int {
	for i := 0; i < len(mapp); i++ {
		for j := 0; j < len(mapp[i]); j++ {
			if mapp[i][j] == "^" {
				return []int{i, j}
			}
		}
	}
	return nil
}

func move(mapp [][]string, location []int) {
	direction := "^"
	end := false
	currentLoc := location

	for !end {

		if direction == "^" {
			currentLoc, direction, end = moveUp(mapp, currentLoc)

		}
		if direction == ">" {
			currentLoc, direction, end = moveRight(mapp, currentLoc)
		}
		if direction == "V" {
			currentLoc, direction, end = moveDown(mapp, currentLoc)
		}
		if direction == "<" {
			currentLoc, direction, end = moveLeft(mapp, currentLoc)
		}
	}
}

func moveLeft(mapp [][]string, location []int) ([]int, string, bool) {
	for i := location[1]; i > 0; i-- {
		if i-1 == 0 {
			if !(mapp[location[0]][i-1] == "#") {
				mapp[location[0]][i] = "X"
				mapp[location[0]][i-1] = "X"
				return []int{location[0], i - 1}, "<", true
			}
		}
		if mapp[location[0]][i-1] == "#" {
			mapp[location[0]][i] = "^"
			return []int{location[0], i}, "^", false
		} else {
			if i == location[1] {
				mapp[location[0]][i-1] = "<"
			} else {
				mapp[location[0]][i-1] = "<"
				mapp[location[0]][i] = "X"
			}
		}

	}
	return []int{}, "<", true

}

func moveRight(mapp [][]string, location []int) ([]int, string, bool) {
	for i := location[1]; i < len(mapp[location[0]]); i++ {
		if i+1 == len(mapp[location[0]]) {
			if !(mapp[location[0]][i+1] == "#") {
				mapp[location[0]][i] = "X"
				mapp[location[0]][i+1] = "X"
				return []int{location[0], i + 1}, ">", true
			}
		}
		if mapp[location[0]][i+1] == "#" {
			mapp[location[0]][i] = "V"
			return []int{location[0], i}, "V", false
		} else {
			if i == location[1] {
				mapp[location[0]][i+1] = ">"
			} else {
				mapp[location[0]][i+1] = ">"
				mapp[location[0]][i] = "X"
			}
		}
	}
	return []int{}, ">", true

}

func moveUp(mapp [][]string, location []int) ([]int, string, bool) {

	for i := location[0]; i > 0; i-- {
		if (i - 1) == 0 {
			if !(mapp[i-1][location[1]] == "#") {
				mapp[i-1][location[1]] = "X"
				mapp[i][location[1]] = "X"
				return []int{i, location[1]}, "^", true
			}
		}
		if mapp[i-1][location[1]] == "#" {
			mapp[i][location[1]] = ">"
			return []int{i, location[1]}, ">", false // true for test
		} else {
			if i == location[0] {
				mapp[i-1][location[1]] = "^"
			} else {
				mapp[i-1][location[1]] = "^"
				mapp[i][location[1]] = "X"
			}
		}
	}
	return []int{}, "^", true
}

func moveDown(mapp [][]string, location []int) ([]int, string, bool) {
	for i := location[0]; i < len(mapp); i++ {
		if (i + 1) == len(mapp) {
			if !(mapp[i+1][location[1]] == "#") {
				mapp[i+1][location[1]] = "X"
				mapp[i][location[1]] = "X"
				return []int{i, location[1]}, "V", true
			}
		}
		if mapp[i+1][location[1]] == "#" {
			mapp[i][location[1]] = "<"
			return []int{i, location[1]}, "<", false // true for test
		} else {
			if i == location[0] {
				mapp[i+1][location[1]] = "V"
			} else {
				mapp[i+1][location[1]] = "V"
				mapp[i][location[1]] = "X"
			}
		}
	}
	return []int{}, "V", true
}

/*
func moveLeft(mapp [][]string, location []int) ([]int, string, bool) {
	for i := location[1]; i > 0; i-- {
		if i-1 == 0 {
			if !(mapp[location[0]][i-1] == "#") {
				mapp[location[0]][i] = "<"
				mapp[location[0]][i-1] = "<"
				return []int{location[0], i - 1}, "<", true
			}
		}
		if mapp[location[0]][i-1] == "#" {
			mapp[location[0]][i] = "^"
			return []int{location[0], i}, "^", false
		} else {
			mapp[location[0]][i-1] = "<"
			mapp[location[0]][i] = "<"
		}

	}
	return []int{}, "<", true

}

func moveRight(mapp [][]string, location []int) ([]int, string, bool) {
	for i := location[1]; i < len(mapp[location[0]]); i++ {
		if i+1 == len(mapp[location[0]]) {
			if !(mapp[location[0]][i+1] == "#") {
				mapp[location[0]][i] = ">"
				mapp[location[0]][i+1] = ">"
				return []int{location[0], i + 1}, ">", true
			}
		}
		if mapp[location[0]][i+1] == "#" {
			mapp[location[0]][i] = "V"
			return []int{location[0], i}, "V", false
		} else {
			mapp[location[0]][i+1] = ">"
			mapp[location[0]][i] = ">"

		}

	}
	return []int{}, ">", true

}

func moveUp(mapp [][]string, location []int) ([]int, string, bool) {

	for i := location[0]; i > 0; i-- {
		if (i - 1) == 0 {
			if !(mapp[i-1][location[1]] == "#") {
				mapp[i-1][location[1]] = "^"
				mapp[i][location[1]] = "^"
				return []int{i, location[1]}, "^", true
			}
		}
		if mapp[i-1][location[1]] == "#" {
			mapp[i][location[1]] = ">"
			return []int{i, location[1]}, ">", false // true for test
		} else {
			mapp[i-1][location[1]] = "^"
			mapp[i][location[1]] = "^"

		}
	}
	return []int{}, "^", true
}

func moveDown(mapp [][]string, location []int) ([]int, string, bool) {
	for i := location[0]; i < len(mapp); i++ {
		if (i + 1) == len(mapp) {
			if !(mapp[i+1][location[1]] == "#") {
				mapp[i+1][location[1]] = "V"
				mapp[i][location[1]] = "V"
				return []int{i, location[1]}, "V", true
			}
		}
		if mapp[i+1][location[1]] == "#" {
			mapp[i][location[1]] = "<"
			return []int{i, location[1]}, "<", false // true for test
		} else {
			mapp[i+1][location[1]] = "V"
			mapp[i][location[1]] = "V"
		}
	}
	return []int{}, "V", true
}
*/
