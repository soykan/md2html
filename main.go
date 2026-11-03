package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func readFile(filePath string) string {
	data, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatal("Error occured while reading file at readFile function.")
	}
	return string(data)
}

func writeFile(filePath string, content string) {
	err := os.WriteFile(filePath, []byte(content), 0644)
	if err != nil {
		log.Fatal("Error occured while writing to file at writeFile function")
	}
}

func convertHeading(line string) string {
	for i := 1; i <= 6; i++ {
		if strings.HasPrefix(line, strings.Repeat("#", i)+" ") {
			return fmt.Sprintf("<h%v>%v</h%v>", i, line[i+1:], i)
		}
	}
	return line
}

func main() {
	if len(os.Args) == 3 {
		mdFilePath := os.Args[1]
		htmlFilePath := os.Args[2]
		fmt.Println(readFile("./deneme.md"))
		fmt.Println(mdFilePath, htmlFilePath)
		fmt.Println(convertHeading("# DENEME"))
	} else {
		log.Fatal("USAGE: ./program <markdown-file-path> <target-html-file-path>")
	}
}
