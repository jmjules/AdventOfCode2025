package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
)

func create_next_day_dir() {
	dir, _ := os.ReadDir(".")

	var day_nums []int
	for _, entry := range dir {

		name_int, err := strconv.Atoi(entry.Name())
		if err != nil {
			continue
		}

		if entry.IsDir() {
			// day_dirs = append(day_dirs, entry)
			day_nums = append(day_nums, name_int)
		}
	}

	if len(day_nums) == 0 {
		fmt.Println("No numbered directories found. Ending script. (Eg.: '01' or '12')")
		return
	}

	slices.Sort(day_nums)
	fmt.Println(day_nums)
	last := day_nums[len(day_nums)-1]
	next := last + 1

	fmt.Printf("Highest day number found was %d.\nCreating folder for day %d:\n", last, next)

	folder_name := fmt.Sprintf("%02d", next)

	os.Mkdir(folder_name, os.ModePerm)

	// add template content
	// main.go
	content := []byte("package main\n\nfunc main() {\n\t\n}\n")
	file_path_main := fmt.Sprintf("%02d/main.go", next)
	err := os.WriteFile(file_path_main, content, os.ModePerm)
	if err != nil {
		panic(err)
	}

	// inputx.txt files
	empty := []byte("")

	file_path_input := fmt.Sprintf("%02d/input.txt", next)
	err = os.WriteFile(file_path_input, empty, os.ModePerm)
	if err != nil {
		panic(err)
	}

	file_path_example := fmt.Sprintf("%02d/input_test.txt", next)
	err = os.WriteFile(file_path_example, empty, os.ModePerm)
	if err != nil {
		panic(err)
	}

}

// TODO: func for specific day folder
// func create_specific_day_dir(i int) {

func main() {
	// TODO: parse flag for specific day
	// switch a) if flag specific b) default next day
	// create_specific_day_dir(i)
	create_next_day_dir()
}
