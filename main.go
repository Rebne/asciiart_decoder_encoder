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

	var outputPath string
	if *writeToOutput {
		if *readInputFromFile && len(args) != 2 {
			fmt.Println("Not a proper amount of arguments")
			return
		} else if *readInputFromFile {
			outputPath = args[1]
		} else {
			outputPath = args[0]
		}
	}

	if *multipleLines || *readInputFromFile {
		var result []string
		if *readInputFromFile {
			if len(args) == 0 {
				fmt.Println("Path to file not inserted")
				return
			} else if len(args) != 1 {
				fmt.Println("Too many paths to file inserted")
				return
			}

			inputPath := args[0]

			result = decodeMultipleLinesFromFile(inputPath, *toEncode)
		} else {
			result = decodeMultipleLines(*toEncode)
		}
		if *writeToOutput {
			writeSliceToFile(&result, outputPath)
		}

		for _, line := range result {
			fmt.Println(line)
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

		var result string
		if *toEncode {
			result = encodeLineArt(lineOfArt)
		} else {
			result = decodeLineArt(lineOfArt)
		}

		if result == "" {
			fmt.Println("Error")
		} else {
			if *writeToOutput {
				writeStringToFile(result, outputPath)
			}
			fmt.Println(result)
		}
	}
}

func writeStringToFile(input string, path string) {
	f, err := os.Create("file.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	_, err = f.WriteString("Hello, World!\n")
	if err != nil {
		log.Fatal(err)
	}
}
func decodeMultipleLinesFromFile(path string, toEncode bool) []string {
	// Opening file with os.Open
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	result := []string{}

	for scanner.Scan() {
		if toEncode {
			result = append(result, encodeLineArt(scanner.Text()))
		} else {
			result = append(result, decodeLineArt(scanner.Text()))
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return result

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
	return len(stack) == 0
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

func decodeMultipleLines(toEncode bool) []string {
	fmt.Println("input text:")
	scanner := bufio.NewScanner(os.Stdin)
	result := []string{}

	for {
		scanner.Scan()
		line := scanner.Text()
		if len(line) == 0 {
			break
		}
		if toEncode {
			result = append(result, encodeLineArt(line))
		} else {
			result = append(result, decodeLineArt(line))
		}
	}

	err := scanner.Err()
	if err != nil {
		log.Fatal(err)
	}

	return result
}

func writeSliceToFile(slice *[]string, path string) {
	f, err := os.Create(path)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	for _, line := range *slice {
		_, err := fmt.Fprintln(f, line)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func displayHelpMessage() {
	fmt.Println("To run the program, use the following command (for example):")
	fmt.Println("go run . [5 #][5 -_]-[5 #]")
	fmt.Println("This displays: #####-_-_-_-_-_-#####")
	fmt.Println()
}
