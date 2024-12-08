package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

/*
type Node struct {
	value    int
	previous *Node
	next     *Node
}

type LinkedList struct {
	head *Node
}

/*
func (list *LinkedList) Insert(key int, value int) {
	newNode := &Node{
		value:    value,
		previous: nil,
		next:     nil,
	}
	//fmt.Println("New node created")

	if list.head == nil {
		//fmt.Println("Head = nil")
		list.head = newNode
		return
	}

	current := list.head
	var foundKey *Node
	//fmt.Println("Key: ", key, " Value:", value)
	for current.value != value && current.next != nil {
		//fmt.Println("No matches for key: ", key, " or value: ", value, " and current.next != nil: ", current.next != nil)
		//fmt.Println("Node.next value: ", current.next)
		//fmt.Println("Next Node")
		if current.value == key {
			foundKey = current
		}
		current = current.next
	}

	if current.value == value {
		if current.previous == nil {
			newNode.next = current
			current.previous = newNode
			list.head = newNode
			return
		}

		newNode.previous = current.previous
		newNode.next = current
		current.previous.next = newNode
		current.previous = newNode

		for current.value != key && current.next != nil {
			current = current.next
		}

		if current.value == key {
			if current.next == nil {
				current.previous.next = nil
				return
			}
			current.next.previous = current.previous
			current.previous.next = current.next
		}

		return
	}

	if foundKey != nil {
		current = foundKey
	}

	if current.value == key {
		//fmt.Println("Found Key: ", key)
		newNode.previous = current
		newNode.next = current.next
		if current.next != nil {
			current.next.previous = newNode
		}
		current.next = newNode

		return
	}

	//fmt.Println(current.value)
	newNode.previous = current
	current.next = newNode
	//fmt.Println("Appended node")
}*/

/*
func (list *LinkedList) PrintList() []int {
	var temp []int
	current := list.head

	for current.next != nil {
		temp = append(temp, current.value)
		current = current.next
	}
	temp = append(temp, current.value)

	return temp
}
*/

/* first attempt
func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	list := &LinkedList{
		head: nil,
	}

	//var sorted [][]int
	sum := 0

	for scanner.Scan() {
		if len(scanner.Text()) != 0 {
			//fmt.Println("second byte: ", scanner.Bytes())
			//fmt.Println("Check: ", scanner.Bytes()[2] == 124)
			if scanner.Bytes()[2] == 124 {
				key, _ := strconv.Atoi(string(scanner.Text()[0:2]))
				value, _ := strconv.Atoi(string(scanner.Text()[3:5]))
				//fmt.Println("Inserting Key: ", key)
				list.Insert(key, key)
				//fmt.Println("Inserting Value: ", value)
				list.Insert(key, value)
			} else if scanner.Bytes()[2] == 44 {
				fmt.Println("Pre Sort: ", line2ints(scanner.Text()))
				sort, _ := list.SortArray(line2ints(scanner.Text()))
				//sorted = append(sorted, sort)
				//fmt.Println(scanner.Text())
				sum = sum + sort[((len(sort)-1)/2)]
			}
			//fmt.Println(scanner.Text())
		}
	}

	fmt.Println("list: ", list.PrintList())
	fmt.Println("Sum: ", sum)
}

func line2ints(line string) []int {
	strs := strings.Split(line, ",")
	//fmt.Println(strs)
	var temp []int
	for i := 0; i < len(strs); i++ {
		num, _ := strconv.Atoi(strs[i])
		temp = append(temp, num)
	}
	return temp
}
*/

// Second approach

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var rules [][]int
	//var sorted [][]int
	sum := 0

	for scanner.Scan() {
		if len(scanner.Text()) != 0 {
			if scanner.Bytes()[2] == 124 {
				var temp []int
				key, _ := strconv.Atoi(string(scanner.Text()[0:2]))
				value, _ := strconv.Atoi(string(scanner.Text()[3:5]))
				temp = append(temp, key, value)
				rules = append(rules, temp)
			} else if scanner.Bytes()[2] == 44 {
				sort := sort(line2ints(scanner.Text()), rules)
				//sorted = append(sorted, sort)
				sum = sum + sort[((len(sort)-1)/2)]
				//fmt.Println(sorted)
			}
		}
	}
	fmt.Println(sum)
}

func line2ints(line string) []int {
	strs := strings.Split(line, ",")
	//fmt.Println(strs)
	var temp []int
	for i := 0; i < len(strs); i++ {
		num, _ := strconv.Atoi(strs[i])
		temp = append(temp, num)
	}
	return temp
}

func sort(line []int, rules [][]int) []int {
	if len(line) < 1 {
		return nil
	}

	var temp []int
	for i := 0; i < len(line); i++ {
		if i != 0 {
			for j := 0; j < len(rules); j++ {
				if rules[j][1] == line[i] {
					check := findIdx(temp, rules[j][0])
					if check == 0 {
						temp = append(temp, line[i])
						break
					}
					insertAtIdx(temp, line[i], check)
				}

			}
		} else {
			temp = append(temp, line[i])
		}
	}
	//fmt.Println("Temp: ", temp)
	return temp
}

func findIdx(line []int, value int) int {
	for i := 0; i < len(line); i++ {
		if line[i] == value {
			return i
		}
	}
	return 0
}

func insertAtIdx(line []int, value int, idx int) []int {
	line = append(line[:idx], append([]int{value}, line[idx:]...)...)
	return line
}
