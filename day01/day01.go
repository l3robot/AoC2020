package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func toInt(list []string) []int {
	var listInt []int
	for i := 0; i < len(list); i++ {
		num, _ := strconv.Atoi(list[i])
		listInt = append(listInt, num)
	}
	return listInt
}

func cAdd(c []int, max int) {
	c[len(c)-1]++
	if c[len(c)-1] == max {
		if !(len(c) == 1) {
			cAdd(c[:len(c)-1], max-1)
			c[len(c)-1] = c[len(c)-2] + 1
		} else {
			panic("You reached the end of the iterator.")
		}
	}
}

func combinations(max int, num int) func() []int {
	var c []int
	for i := 0; i < num; i++ {
		c = append(c, i)
	}
	return func() []int {
		cAdd(c, max)
		return c
	}
}

func part1(list []int) {
	// Defines the combinations iterator.
	cIterator := combinations(len(list), 2)
	for {
		c := cIterator()
		sum := list[c[0]] + list[c[1]]
		if sum == 2020 {
			fmt.Println("Part One")
			fmt.Println("--------")
			fmt.Println(list[c[0]], list[c[1]])
			fmt.Println(list[c[0]] * list[c[1]])
			fmt.Println("")
			break
		}
	}
}

func part2(list []int) {
	// Defines the combinations iterator.
	cIterator := combinations(len(list), 3)
	for {
		c := cIterator()
		sum := list[c[0]] + list[c[1]] + list[c[2]]
		if sum == 2020 {
			fmt.Println("Part Two")
			fmt.Println("--------")
			fmt.Println(list[c[0]], list[c[1]], list[c[2]])
			fmt.Println(list[c[0]] * list[c[1]] * list[c[2]])
			fmt.Println("")
			break
		}
	}
}

func main() {
	// Read the file.
	bytes, err := ioutil.ReadFile("day01/data.txt")
	check(err)

	// Split lines of the file.
	data := string(bytes)
	dataList := strings.Split(data, "\n")

	// Convert list items to int.
	dataListInt := toInt(dataList)

	// Part One
	part1(dataListInt)

	// Part Two
	part2(dataListInt)
}
