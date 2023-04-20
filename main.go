package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"

	"github.com/hculpan/kmonkey/lexer"
	"github.com/hculpan/kmonkey/token"
)

func main() {
	//	genTestFile()

	lines, err := ReadFile("test.mky")

	if err != nil {
		fmt.Println(err)
		return
	}

	l := lexer.NewLexer(lines)

	for {
		t := l.NextToken()
		fmt.Printf("%s: '%s' at %d:%d\n", t.Type, t.Literal, t.Line, t.Pos)
		if t.Type == token.EOF {
			break
		}
	}
}

// Write a function to read a text file and return it as a slice of strings
func ReadFile(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}

func genTestFile() {
	f, err := os.OpenFile("long_text.tokens", os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	tokens := []string{"+", ",", "(", ")", "let", "add", "fn", "a", ";", "*", "abcdefg_123", "123", "456"}
	for i := 0; i < 5000; i++ {
		n := rand.Intn(50)
		line := ""
		for j := 0; j < n; j++ {
			idx := rand.Intn(len(tokens))
			line += tokens[idx] + " "
		}
		f.WriteString(line + "\n")
	}
}
