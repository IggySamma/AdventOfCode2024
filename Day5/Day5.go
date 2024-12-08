package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
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
	sumP2 := 0

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
				temp := line2ints(scanner.Text())
				check := compare(sort, temp)
				sum = sum + compare(sort, temp)
				if check == 0 {
					sumP2 = sumP2 + sort[(len(sort)-1)/2]
				}
				//sorted = append(sorted, sort)

				//fmt.Println(sorted)
			}
		}
	}
	fmt.Println("Part 1: ", sum)
	fmt.Println("Part 2: ", sumP2)
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

func compare(sort []int, temp []int) int {
	for i := 0; i < len(sort); i++ {
		if sort[i] != temp[i] {
			return 0
		}
	}
	return sort[(len(sort)-1)/2]
}

func sort(line []int, rules [][]int) []int {
	if len(line) < 1 {
		return nil
	}

	slices.SortFunc(line, func(a, b int) int {
		for _, rule := range rules {
			if rule[0] == a && rule[1] == b {
				return -1
			}
			if rule[0] == b && rule[1] == a {

				return 1
			}
		}
		if a < b {
			return -1
		}
		if a > b {
			return 1
		}
		return 0
	})

	return line
}
