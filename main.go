package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var expressionForCheck string
var expressionForDecoding string
var regexForCheck *regexp.Regexp
var regexForDecoding *regexp.Regexp

func init() {
	expressionForCheck = `\[[^\d]+ [^\[]+\]|\[\d \]|\[\d[^\s\]]*\]`
	regexForCheck = regexp.MustCompile(expressionForCheck)

	expressionForDecoding = `\[\d [^\[]+\]|.`
	regexForDecoding = regexp.MustCompile(expressionForDecoding)
}

func main() {
	// Brackets balande, first item digit, space between, some value after space
	multipleLines := flag.Bool("m", false, "Enable to enter multiple lines of input")
	toEncode := flag.Bool("e", false, "Select encoding mode for input line")
	writeToOutput := flag.Bool("o", false, "Write output to specified file")
	readInputFromFile := flag.Bool("i", false, "Read input from a file")
	readFromFileAndWriteToFile := flag.Bool("io", false, "Read input from file & write this to output file")

	flag.Parse()

	args := flag.Args()

	if *readFromFileAndWriteToFile {
		*readInputFromFile = true
		*writeToOutput = true
	}

	if *multipleLines || *readInputFromFile {
		if *readInputFromFile {
			if len(args) == 0 {
				fmt.Println("Path to file not inserted")
				return
			} else if len(args) != 1 {
				fmt.Println("Too many paths to file inserted")
				return
			}

			path := args[0]

			decodeMultipleLinesFromFile(&result, path, *toEncode)
		} else {
			decodeMultipleLines(&result, *toEncode)
		}
	} else {

		if len(args) == 0 {
			displayHelpMessage()
			return
		} else if len(args) != 1 {
			fmt.Println("The program only accepts one input. You entered too many.")
			return
		}

		lineOfArt := args[0]
		if *toEncode {
			result := encodeLineArt(lineOfArt)
		} else {
			result := decodeLineArt(lineOfArt)
		}

		if result == "" {
			fmt.Println("Error")
		} else {
			if *writeToOutput {
				return
			}
			fmt.Println(result)
		}
	}
}

func decodeMultipleLinesFromFile(result *string, path string, toEncode bool) {
	// Opening file with os.Open
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		if toEncode {
			*result += encodeLineArt(scanner.Text()) + "\n"
		} else {
			*result += decodeLineArt(scanner.Text()) + "\n"
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

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

func encodeLineArt(input string) string {
	var resultArray []string
	length := len(input)

	for i := 0; i < length; i++ {
		if i+1 < length && input[i] == input[i+1] {
			resultArray = append(resultArray, string(input[i]))
		} else if i+2 < length && input[i] == input[i+2] {
			resultArray = append(resultArray, input[i:i+2])
			i++
		} else {
			resultArray = append(resultArray, string(input[i]))
		}
	}

	var result string
	var count int
	i := 0
	length = len(resultArray)
	for i < len(resultArray) {
		if i+1 < length && resultArray[i] == resultArray[i+1] {
			count++
		} else {
			if count >= 3 {
				result += fmt.Sprintf(`[%d %s]`, count+1, resultArray[i])
				count = 0
			} else if count > 0 {
				for j := 0; j <= count; j++ {
					result += resultArray[i]
				}
				count = 0
			} else {
				result += resultArray[i]
			}
		}
		i++
	}
	return result
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

func decodeMultipleLines(result *string, toEncode bool) {
	fmt.Println("input text:")
	scanner := bufio.NewScanner(os.Stdin)
	for {
		scanner.Scan()
		line := scanner.Text()
		if len(line) == 0 {
			break
		}
		if toEncode {
			*result += encodeLineArt(line) + "\n"
		} else {
			*result += decodeLineArt(line) + "\n"
		}
	}

	err := scanner.Err()
	if err != nil {
		log.Fatal(err)
	}
}

func displayHelpMessage() {
	fmt.Println("To run the program, use the following command (for example):")
	fmt.Println("go run . [5 #][5 -_]-[5 #]")
	fmt.Println("This displays: #####-_-_-_-_-_-#####")
	fmt.Println()
}
