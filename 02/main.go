package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func is_fake_id_p1(id_string string) bool {
	if len(id_string)%2 != 0 {
		return false
	}

	mid := len(id_string) / 2

	part1 := id_string[:mid]
	part2 := id_string[mid:]

	if part1 == part2 {
		return true
	} else {
		return false
	}

}
func is_fake_id_p2(id_string string) bool {

	full := len(id_string)
	mid := len(id_string) / 2

	for size := 1; size <= mid; size++ {

		if full%size != 0 {
			continue
		}

		token := id_string[:size]

		matching := true
		for i := 0; i < full; i += size {
			subs := id_string[i : i+size]
			if token != subs {
				matching = false
				break
			}
		}

		if matching {
			return true
		}
	}

	return false

}

func main() {

	file, _ := os.Open("input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	line := scanner.Text()

	id_ranges := strings.Split(line, ",")

	sum1 := 0
	sum2 := 0
	for _, id_range := range id_ranges {
		// fmt.Println(id_range)
		parts := strings.Split(id_range, "-")

		range_start, _ := strconv.Atoi(parts[0])
		range_end, _ := strconv.Atoi(parts[1])

		for i := range_start; i <= range_end; i++ {

			i_str := strconv.Itoa(i)

			if is_fake_id_p1(i_str) {
				sum1 += i
				// println("-->Found Fake Id:", i)
			}
			if is_fake_id_p2(i_str) {
				sum2 += i
				// println("-->Found Fake Id:", i)
			}

		}

	}

	fmt.Println("Part 1: Final sum:", sum1)
	fmt.Println("Part 2: Final sum:", sum2)
}
