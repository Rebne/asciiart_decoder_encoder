package main

import (
	"fmt"
	"html/template"
	"net/http"
	"regexp"
	"strconv"
	"strings"
)

var expressionForCheck string
var expressionForDecoding string
var regexForCheck *regexp.Regexp
var regexForDecoding *regexp.Regexp

var templates map[string]*template.Template

func init() {
	expressionForCheck = `\[[^\d]+ [^\[]+\]|\[\d \]|\[\d[^\s\]]*\]`
	regexForCheck = regexp.MustCompile(expressionForCheck)

	expressionForDecoding = `\[\d+ [^\[]+\]|.`
	regexForDecoding = regexp.MustCompile(expressionForDecoding)

	if templates == nil {
		templates = make(map[string]*template.Template)
	}
	templates["index"] = template.Must(template.ParseFiles("template/index.html", "template/base.html"))
	templates["decode"] = template.Must(template.ParseFiles("template/decoder.html", "template/base.html"))
}

func main() {

	http.HandleFunc("/", getIndex)
	http.HandleFunc("/method", chooseMethod)
	http.ListenAndServe(":8088", nil)
}

func chooseMethod(w http.ResponseWriter, r *http.Request) {
	method := r.FormValue("processMethod")

	switch method {
	case "decode":
		decodePage(w, r)
	case "encode":
		encodePage(w, r)
	default:
		http.Error(w, "Invalid action", http.StatusBadRequest)
	}
}

func getIndex(w http.ResponseWriter, r *http.Request) {
	var currData Data
	currData.StatusCode = http.StatusOK
	renderTemplate(w, "index", "base", currData)
}

func decodePage(w http.ResponseWriter, r *http.Request) {
	var currData Data
	r.ParseForm()
	input := r.PostFormValue("input")

	currData.Array = decodeMultipleLines(false, input)
	// Handling malformed string and returning status BadRequest
	if len(currData.Array) == 0 {
		currData.StatusCode = http.StatusBadRequest
		renderTemplate(w, "index", "base", currData)
	} else {
		currData.StatusCode = http.StatusAccepted

		// for _, line := range result {
		// 	fmt.Println(line)
		// }
		renderTemplate(w, "decode", "base", currData)
	}
}

func encodePage(w http.ResponseWriter, r *http.Request) {
	var currData Data
	r.ParseForm()
	input := r.PostFormValue("input")

	currData.Array = decodeMultipleLines(true, input)
	currData.StatusCode = http.StatusAccepted

	// for _, line := range result {
	// 	fmt.Println(line)
	// }
	renderTemplate(w, "decode", "base", currData)
}

type Data struct {
	Array      []string
	StatusCode int
}

func renderTemplate(w http.ResponseWriter, name string, template string, viewModel interface{}) {
	tmpl, ok := templates[name]
	if !ok {
		http.Error(w, "The template does not exist", http.StatusInternalServerError)
	}
	err := tmpl.ExecuteTemplate(w, template, viewModel)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func encodeLine(input string) string {
	var resultArray []string
	length := len(input)

	// Splitting the input to slice/array for reworking into encoded version later.
	for i := 0; i < length; i++ {
		if i+1 < length && input[i] == input[i+1] {
			resultArray = append(resultArray, string(input[i]))
		} else if i+2 < length && input[i] == input[i+2] ||
			len(resultArray) > 0 && len(resultArray[len(resultArray)-1]) != 1 &&
				input[i:i+2] == resultArray[len(resultArray)-1] {
			resultArray = append(resultArray, input[i:i+2])
			i++
		} else {
			resultArray = append(resultArray, string(input[i]))
		}
	}

	// Constructing result string from the array by counting consecutive elements
	// and encoding if there are at least 4 consecutive elements
	var result string
	var count int
	i := 0
	length = len(resultArray)
	for i < len(resultArray) {
		if i+1 < length && resultArray[i] == resultArray[i+1] {
			count++
		} else {
			if count > 0 {
				result += fmt.Sprintf(`[%d %s]`, count+1, resultArray[i])
				count = 0
			} else {
				result += resultArray[i]
			}
		}
		i++
	}
	return result
}

func decodeLine(input string) string {
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

func decodeMultipleLines(e bool, s string) []string {

	array := strings.Split(s, "\n")

	if e {
		for i, line := range array {
			if array[i] == "" {
				return nil
			}
			array[i] = encodeLine(line)
		}
	} else {
		for i, line := range array {
			// checking for error
			if array[i] == "" {
				return nil
			}
			array[i] = decodeLine(line)
		}
	}

	return array
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
