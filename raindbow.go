package main

import "fmt"

const (
	reset = "\x1b[0m"
)

func addColorToText(arr []string) []string {
	fmt.Println("1 Red")
	fmt.Println("2 Green")
	fmt.Println("3 Blue")
	fmt.Println("4 Rainbow")
	fmt.Print("Insert number for color: ")

	var color string
	_, err := fmt.Scan(&color)

	if err != nil {
		fmt.Println("Error reading input:", err)
		return []string{}
	}

	switch color {
	case "1":
		return paintStringToColor(arr, color)
	case "2":
		return paintStringToColor(arr, color)
	case "3":
		return paintStringToColor(arr, color)
	case "4":
		return
	}

}

func paintStringToColor(arr []string, key string) []string {
	colors := map[string]string{
		"1": "\x1b[91m", // BrightRed
		"2": "\x1b[92m", // BrightGreen
		"3": "\x1b[94m", // BrightBlue
	}
	for i := range arr {
		arr[i] = colors[key] + arr[i] + reset
	}

	return arr
}
func paintAsRainbow(asciiArt []string) []string {
	rainbowColors := map[int]string{
		1: "\x1b[31m",             // Red
		2: "\x1b[33m",             // Yellow
		3: "\x1b[38;2;255;165;0m", // Orange
		4: "\x1b[32m",             // Green
		5: "\x1b[36m",             // Cyan
		6: "\x1b[34m",             // Blue
		7: "\x1b[35m",             // Magenta
	}
	// Print the ANSI code values for each color

	longestRow := 0
	result := [][]string{}
	for i := range asciiArt {
		result = append(result, make([]string, 0))
		for _, char := range asciiArt[i] {
			result[i] = append(result[i], string(char))
		}
		if len(result[i]) > len(result[longestRow]) {
			longestRow = i
		}
	}

	// color := 0
	// var full string
	// for i := range result {
	// 	for _, val := range result[i] {
	// 		full += val
	// 	}
	// 	full = rainbowColors[color+1] + full + reset
	// 	fmt.Println(full)
	// 	full = ""
	// 	color = (color + 1) % 7

	// }

	// for range result {
	// 	fmt.Println()
	// }

	color := 0

	for idx := range result[longestRow] {
		row, col := longestRow, idx
		for row >= 0 {
			if col < len(result[row]) && result[row][col] == " " {
				row--
				continue
			}
			if col < len(result[row]) {
				result[row][col] = rainbowColors[color+1] + result[row][col] + reset
			}
			row--
		}
		row = longestRow + 1
		for row < len(result) {
			if result[row][col] == " " {
				row++
				continue
			}
			if col < len(result[row]) {
				result[col][row] = rainbowColors[color+1] + result[row][col] + reset
			}
			row++
		}
		color = (color + 1) % 7
	}
	var tmp string
	for i := range result {
		for _, val := range result[i] {
			tmp += val
		}
		fmt.Println(tmp)
		tmp = ""
	}

}
