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
		return nil
	}

	switch color {
	case "1":
		return paintStringToColor(arr, color)
	case "2":
		return paintStringToColor(arr, color)
	case "3":
		return paintStringToColor(arr, color)
	case "4":
		return paintAsRainbow(arr)
	default:
		fmt.Println("Invalid number inserted. No coloring added.")
		return nil

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

	longestRow := 0
	charArray := [][]string{}
	for i := range asciiArt {
		charArray = append(charArray, make([]string, 0))
		for _, char := range asciiArt[i] {
			charArray[i] = append(charArray[i], string(char))
		}
		if len(charArray[i]) > len(charArray[longestRow]) {
			longestRow = i
		}
	}

	color := 0

	for idx := range charArray[longestRow] {
		row, col := longestRow, idx
		for row >= 0 {
			if col < len(charArray[row]) && charArray[row][col] == " " {
				row--
				continue
			}
			if col < len(charArray[row]) {
				charArray[row][col] = rainbowColors[color+1] + charArray[row][col] + reset
			}
			row--
		}
		row = longestRow + 1
		for row < len(charArray) {
			if col < len(charArray[row]) && charArray[row][col] == " " {
				row++
				continue
			}
			if col < len(charArray[row]) {
				charArray[row][col] = rainbowColors[color+1] + charArray[row][col] + reset
			}
			row++
		}
		color = (color + 1) % 7
	}
	result := []string{}
	var tmp string
	for i := range charArray {
		for _, val := range charArray[i] {
			tmp += val
		}
		result = append(result, tmp)
		tmp = ""
	}
	return result

}
