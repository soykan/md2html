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

func convertLineBreak(line string) string {
	if line[len(line)-2:] == "  " {
		return fmt.Sprintf("%v<br/>", line[:len(line)-2])
	}
	return line
}

func convertEmptyLine(line string) string {
	if strings.TrimSpace(line) == "" {
		return "<br/>"
	}
	return line
}

func convertBoldText(text string) string {
	if (text[:2] == "**" && text[len(text)-2:] == "**") ||
		(text[:2] == "__" && text[len(text)-2:] == "__") {
		return fmt.Sprintf("<b>%v</b>", text[2:len(text)-2])
	}
	return text
}

func convertItalicText(text string) string {
	if (text[:1] == "*" && text[len(text)-1:] == "*") ||
		(text[:1] == "_" && text[len(text)-1:] == "_") {
		return fmt.Sprintf("<i>%v</i>", text[1:len(text)-1])
	}
	return text
}

func main() {
	if len(os.Args) == 3 {
		mdFilePath := os.Args[1]
		htmlFilePath := os.Args[2]
		fmt.Println(readFile("./development.md"))
		fmt.Println(mdFilePath, htmlFilePath)
		fmt.Println(convertHeading("# DENEME"))
		fmt.Println(convertLineBreak("deneme   "))
		fmt.Println(convertEmptyLine("    "))
		fmt.Println(convertBoldText("__deneme 123__"))
		fmt.Println(convertItalicText("*deneme 123*"))
	} else {
		log.Fatal("USAGE: ./program <markdown-file-path> <target-html-file-path>")
	}
}
