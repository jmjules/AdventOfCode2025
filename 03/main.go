package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// memoization?
func next_largest_digit_in_string(line string, digit_len int, current_num string) string {
	//base case?
	if len(line) == 2 {
		return string(line[0]) //oder line[1]
		// return current_num
	}
	if digit_len-1 == 0 {

	}
	///hmmmmm?????

	contender := 0
	contender_idx := -1

	for i, char := range line[:len(line)-(digit_len-1)] {
		digit := int(char - '0')
		if digit > contender {
			contender = digit
			contender_idx = i

			if contender == 9 {
				break
			}
		}
	}

	// get values for current one
	// go deeper for next values
	// add together and return? == profit?

}

// func check_substring

func main() {
	file, _ := os.Open("input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	sum := 0

	for scanner.Scan() {
		line := scanner.Text()
		wordlength := len(line)

		contender := 0
		contender_idx := -1
		for i, char := range line[:wordlength-1] {

			digit := int(char - '0')
			if digit > contender {
				contender = digit
				contender_idx = i
				// fmt.Println(rune(contender), ",", contender_idx)

				if contender == 9 {
					break
				}
			}
		}

		contender_2 := 0
		// contender_2_idx := -1
		// fmt.Print("2nd contenders: ")

		for _, char := range line[contender_idx+1:] {

			digit := int(char - '0')
			// fmt.Print(digit)
			if digit > contender_2 {
				contender_2 = digit
				// contender_idx = i
				// fmt.Println(rune(contender))

				if contender_2 == 9 {
					break
				}
			}
		}
		// fmt.Println()

		// fmt.Println("1st/2nd:", contender, contender_2)

		joltage, _ := strconv.Atoi(strconv.Itoa(contender) + strconv.Itoa(contender_2))
		// fmt.Println(joltage)

		sum += joltage
		// fmt.Println("----------")
	}

	fmt.Println("Sum is:", sum)
}
