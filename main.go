package main

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

const (
	CELL_NUMBER = 30000000
)

const (
	PTUP = iota
	PTDOWN
	VALUEUP
	VALUEDOWN
	READOUT
	READIN
	BEGINLOOP
	ENDLOOP
)

type OptCode struct {
	Type  int
	Value int
}

const tab = "\t"

func main() {
	var inputFile, outputFile string

	flag.StringVar(&inputFile, "file", "", "path to program file")
	flag.StringVar(&outputFile, "out", "", "path to output file")
	flag.Parse()

	if len(inputFile) == 0 {
		flag.Usage()
		os.Exit(1)
	}

	buf, err := ioutil.ReadFile(inputFile)
	if err != nil {
		log.Fatal(err.Error())
	}
	program := string(buf)
	programLen := len(program)
	optCodes := make([]*OptCode, 0)
	last := &OptCode{
		Type: -1,
	}
	hasReadin := false
	hasOutput := false
	for i := 0; i < programLen; i++ {
		if len(optCodes) > 0 {
			last = optCodes[len(optCodes)-1]
			optCodes = optCodes[:len(optCodes)-1]
		}
		switch program[i] {
		case '>':
			if last.Type == PTUP {
				last.Value++
				optCodes = append(optCodes, last)
			} else {
				if last.Type != -1 {
					optCodes = append(optCodes, last)
				}
				optCodes = append(optCodes, &OptCode{
					Type:  PTUP,
					Value: 1,
				})
			}
		case '<':
			if last.Type == PTDOWN {
				last.Value++
				optCodes = append(optCodes, last)
			} else {
				if last.Type != -1 {
					optCodes = append(optCodes, last)
				}
				optCodes = append(optCodes, &OptCode{
					Type:  PTDOWN,
					Value: 1,
				})
			}
		case '+':
			if last.Type == VALUEUP {
				last.Value++
				optCodes = append(optCodes, last)
			} else {
				if last.Type != -1 {
					optCodes = append(optCodes, last)
				}
				optCodes = append(optCodes, &OptCode{
					Type:  VALUEUP,
					Value: 1,
				})
			}
		case '-':
			if last.Type == VALUEDOWN {
				last.Value++
				optCodes = append(optCodes, last)
			} else {
				if last.Type != -1 {
					optCodes = append(optCodes, last)
				}
				optCodes = append(optCodes, &OptCode{
					Type:  VALUEDOWN,
					Value: 1,
				})
			}
		case '.':
			if last.Type != -1 {
				optCodes = append(optCodes, last)
			}
			optCodes = append(optCodes, &OptCode{
				Type: READOUT,
			})
			hasOutput = true
		case ',':
			if last.Type != -1 {
				optCodes = append(optCodes, last)
			}
			optCodes = append(optCodes, &OptCode{
				Type: READIN,
			})
			hasReadin = true
		case '[':
			if last.Type != -1 {
				optCodes = append(optCodes, last)
			}
			optCodes = append(optCodes, &OptCode{
				Type: BEGINLOOP,
			})
		case ']':
			if last.Type != -1 {
				optCodes = append(optCodes, last)
			}
			optCodes = append(optCodes, &OptCode{
				Type: ENDLOOP,
			})
		default:
			optCodes = append(optCodes, last)
			continue
		}
	}

	// Print out go code!
	sb := &bytes.Buffer{}
	tabs := 0
	addLine(sb, "package main", tabs)
	if hasOutput {
		addLine(sb, "import \"fmt\"", tabs)
	}
	if hasReadin {
		addLine(sb, "import \"os\"", tabs)
	}
	addLine(sb, "func main() {", tabs)
	tabs++
	addLine(sb, fmt.Sprintf("buffer := make([]byte, %d)", CELL_NUMBER), tabs)
	addLine(sb, "ptr := 0", tabs)
	if hasReadin {
		addLine(sb, "b := make([]byte, 1)", tabs)
	}

	for _, optCode := range optCodes {
		switch optCode.Type {
		case PTUP:
			addLine(sb, fmt.Sprintf("ptr = ptr + %d", optCode.Value), tabs)

		case PTDOWN:
			addLine(sb, fmt.Sprintf("ptr = ptr - %d", optCode.Value), tabs)

		case VALUEUP:
			addLine(sb, fmt.Sprintf("buffer[ptr] = buffer[ptr] + byte(%d)", optCode.Value), tabs)

		case VALUEDOWN:
			addLine(sb, fmt.Sprintf("buffer[ptr] = buffer[ptr] - byte(%d)", optCode.Value), tabs)

		case READOUT:
			addLine(sb, "fmt.Print(string(buffer[ptr]))", tabs)

		case READIN:
			addLine(sb, "os.Stdin.Read(b)", tabs)
			addLine(sb, "buffer[ptr] = b[0]", tabs)

		case BEGINLOOP:
			addLine(sb, "for buffer[ptr] != 0 {", tabs)
			tabs++

		case ENDLOOP:
			tabs--
			addLine(sb, "}", tabs)
		}
	}
	tabs--
	addLine(sb, "}", tabs)

	err = ioutil.WriteFile(outputFile, sb.Bytes(), 0644)
	if err != nil {
		fmt.Println(err.Error())
	}

}

func addLine(sb *bytes.Buffer, value string, tabs int) {
	for i := 0; i < tabs; i++ {
		sb.WriteString("\t")
	}
	sb.WriteString(value)
	sb.WriteString("\n")

}
