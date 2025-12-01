package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const START_INDEX = 50
const MIN_INDEX = 0
const MAX_INDEX = 99

func main() {

	file, _ := os.Open("input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	current_idx := START_INDEX
	counter_ending := 0
	counter_passing := 0

	for scanner.Scan() {
		line := scanner.Text()

		direction := rune(line[0])
		amount, _ := strconv.Atoi(line[1:])
		rel_amount := amount % 100

		// multiplier := 1
		// if direction == 'L' {
		// 	multiplier = -1
		// }

		fmt.Println("-----------")

		if direction == 'L' {
			delta := current_idx - rel_amount
			if delta < 0 {
				if delta-100 != 0 && current_idx != 0 {
					counter_passing++
					println("passing")
				}
				current_idx = 100 + delta
			} else {
				current_idx = current_idx - rel_amount
			}

		} else {
			delta := current_idx + rel_amount

			if delta > 99 {
				if delta-100 != 0 {
					counter_passing++
					println("passing")
				}
				current_idx = delta - 100
			} else {
				current_idx = current_idx + rel_amount
			}
		}

		if current_idx == 0 {
			counter_ending++
		}

		if amount/100 != 0 {
			counter_passing += amount / 100
			fmt.Println("ecnountered", amount, "adding", amount/100)
		}

		fmt.Println(string(direction), amount, rel_amount)
		fmt.Println(current_idx)

	}

	fmt.Println("Part 1: Final count:", counter_ending)
	fmt.Println("Part 2: Final count:", counter_ending+counter_passing)

}
