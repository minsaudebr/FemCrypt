package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strconv"
)

func obfuscateCode(code string) string {
	var result string
	scanner := bufio.NewScanner(bufio.NewReaderString(code))
	variableMap := make(map[string]string)
	functionMap := make(map[string]string)
	var varCounter, funcCounter int

	for scanner.Scan() {
		line := scanner.Text()

		// Obfuscate variables (naive approach)
		varRegex := regexp.MustCompile(`\bvar\s+(\w+)`)
		line = varRegex.ReplaceAllStringFunc(line, func(m string) string {
			varName := string(m[len("var "):]) // Extract variable name
			if _, ok := variableMap[varName]; !ok {
				varCounter++
				obfuscatedName := "v" + strconv.Itoa(varCounter)
				variableMap[varName] = obfuscatedName
			}
			return "var " + variableMap[varName]
		})

		// Obfuscate function names (naive approach)
		funcRegex := regexp.MustCompile(`\bfunc\s+(\w+)`)
		line = funcRegex.ReplaceAllStringFunc(line, func(m string) string {
			funcName := string(m[len("func "):]) // Extract function name
			if _, ok := functionMap[funcName]; !ok {
				funcCounter++
				obfuscatedName := "f" + strconv.Itoa(funcCounter)
				functionMap[funcName] = obfuscatedName
			}
			return "func " + functionMap[funcName]
		})

		result += line + "\n"
	}

	return result
}

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: go run obfuscator.go <input_file.go> <output_file.go>")
		return
	}

	inputFile := os.Args[1]
	outputFile := os.Args[2]

	// Read the input Go code
	code, err := ioutil.ReadFile(inputFile)
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return
	}

	// Obfuscate the code
	obfuscatedCode := obfuscateCode(string(code))

	// Write the obfuscated code to output file
	err = ioutil.WriteFile(outputFile, []byte(obfuscatedCode), 0644)
	if err != nil {
		fmt.Printf("Error writing file: %v\n", err)
		return
	}

	fmt.Println("Code obfuscation completed successfully!")
}
