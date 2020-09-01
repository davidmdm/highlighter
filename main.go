package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	words := os.Args[1:]

	for {
		line, err := reader.ReadString('\n')

		if err != nil {
			if err == io.EOF {
				break
			}
			fatalf("failed to read next line: %v", err)
		}
		for _, word := range words {
			line = highlightWord(line, word)
		}
		fmt.Print(line)
	}
}

func highlightWord(line, word string) string {
	split := strings.Split(line, word)
	return strings.Join(split, "\033[47;30;1m"+word+"\033[0m")
}

func fatalf(format string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, format+"\n", args...)
	os.Exit(1)
}
