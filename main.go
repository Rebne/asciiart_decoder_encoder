package main

import (
	"flag"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

var expressionForCheck string
var expressionForDecoding string
var regexForCheck *regexp.Regexp
var regexForDecoding *regexp.Regexp

func init() {
	expressionForCheck = `\[[^\d] [^\[]+\]|\[\d \]|\[\d[^ ].*[^\[]\]`
	regexForCheck = regexp.MustCompile(expressionForCheck)

	expressionForDecoding = `\[\d [^\[]+\]|.`
	regexForDecoding = regexp.MustCompile(expressionForDecoding)
}

func checkForBalancedBrackets(input string) bool {
	stack := []rune{}

	for _, char := range input {
		if char == '[' {
			stack = append(stack, char)
		} else if char == ']' {
			if len(stack) == 0 || stack[len(stack)-1] != '[' {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}
	if len(stack) == 0 {
		return true
	}
	return false
}

func isValidLineArt(input string) bool {
	if !checkForBalancedBrackets(input) {
		return false
	}
	if regexForCheck.MatchString(input) {
		return false
	}
	return true
}

func decodeLineArt(input string) string {
	matches := regexForDecoding.FindAllString(input, -1)
	var result string
	for _, match := range matches {
		if match[0] == '[' {
			if !isValidLineArt(match) {
				return ""
			}

			var num int
			index := 0
			for match[index] != ' ' {
				index++
			}
			num, _ = strconv.Atoi(match[1:index])
			result += strings.Repeat(match[index+1:len(match)-1], num)
		} else {
			result += match
		}
	}
	return result

}

func main() {
	// Brackets balande, first item digit, space between, some value after space
	flag.Parse()
	args := flag.Args()
	fmt.Println(args)

	displayHelpMessage := func() {
		fmt.Println("To run the program, use the following command (for example):")
		fmt.Println("go run . [5 #][5 -_]-[5 #]")
		fmt.Println("This displays: #####-_-_-_-_-_-#####")
		fmt.Println()
	}
	if len(args) == 0 {
		displayHelpMessage()
		return
	} else if len(args) != 1 {
		fmt.Println("The program only accepts one input. You entered too many.")
		return
	}

	lineOfArt := flag.Args()[0]

	result := decodeLineArt(lineOfArt)

	if result == "" {
		fmt.Println("Error")
	} else {
		fmt.Println(result)
	}
}
