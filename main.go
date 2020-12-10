// https://adventofcode.com/2020/day/8
package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func contains(items []int, item int) bool {
	for _, anItem := range items {
		if anItem == item {
			return true
		}
	}

	return false
}

func operate(lines []string) (int, bool) {
	var infinite = false
	var acc = 0
	var executedLines []int

	for i := 0; i < len(lines); i++ {
		if contains(executedLines, i) {
			infinite = true
			break
		}

		executedLines = append(executedLines, i)
		line := lines[i]

		if strings.HasPrefix(line, "nop ") {
			continue
		}

		if strings.HasPrefix(line, "jmp ") {
			value, _ := strconv.Atoi(line[4:])
			i += value - 1
			continue
		}

		if strings.HasPrefix(line, "acc ") {
			value, _ := strconv.Atoi(line[4:])
			acc += value
			continue
		}
	}

	return acc, infinite
}

func main() {
	fc, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Println(err)
		return
	}

	lines := strings.Split(string(fc), "\n")
	acc, _ := operate(lines)

	fmt.Println(acc)

	// Part 2
	// Let's recreate the list of operations by replacing exactly 1 nop to jmp and vice versa
	for idx, line := range lines {
		if strings.HasPrefix(line, "nop ") || strings.HasPrefix(line, "jmp ") {
			ops := make([]string, len(lines))
			copy(ops, lines)

			if strings.HasPrefix(line, "nop ") {
				ops[idx] = strings.ReplaceAll(ops[idx], "nop ", "jmp ")
			} else if strings.HasPrefix(line, "jmp ") {
				ops[idx] = strings.ReplaceAll(ops[idx], "jmp ", "nop ")
			}

			acc, infinite := operate(ops)
			if infinite {
				continue
			}

			fmt.Println(acc)
			break
		}
	}
}