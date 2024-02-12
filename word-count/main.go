package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {

	var wordCount int
	var words []string
	if len(os.Args) >= 2 {
		fileContents, err := os.ReadFile(os.Args[1])
		if err == nil {
			words = strings.Fields(string(fileContents))
			wordCount = len(words)
			fmt.Println("Found ", wordCount, "ocurrences")
		} else {
			log.Panicln("File doesn`t exist")
		}
/	}
}
