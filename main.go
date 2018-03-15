package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
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
	SETZERO
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
	optCodes := NewFifo()
	last := &OptCode{
		Type: -1,
	}
	hasReadin := false
	hasOutput := false
	for i := 0; i < programLen; i++ {
		if optCodes.HasNext() {
			last = optCodes.PopLast()
		}
		switch program[i] {
		case '>':
			if last.Type == PTUP {
				last.Value++
			} else {
				if last.Type != -1 {
					optCodes.Push(last)
				}
				optCodes.Push(&OptCode{
					Type:  PTUP,
					Value: 1,
				})
			}
		case '<':
			if last.Type == PTDOWN {
				last.Value++
			} else {
				if last.Type != -1 {
					optCodes.Push(last)
				}
				optCodes.Push(&OptCode{
					Type:  PTDOWN,
					Value: 1,
				})
			}
		case '+':
			if last.Type == VALUEUP {
				last.Value++
			} else {
				if last.Type != -1 {
					optCodes.Push(last)
				}
				optCodes.Push(&OptCode{
					Type:  VALUEUP,
					Value: 1,
				})
			}
		case '-':
			if last.Type == VALUEDOWN {
				last.Value++
			} else {
				if last.Type != -1 {
					optCodes.Push(last)
				}
				optCodes.Push(&OptCode{
					Type:  VALUEDOWN,
					Value: 1,
				})
			}
		case '.':
			if last.Type != -1 {
				optCodes.Push(last)
			}
			optCodes.Push(&OptCode{
				Type: READOUT,
			})
			hasOutput = true
		case ',':
			if last.Type != -1 {
				optCodes.Push(last)
			}
			optCodes.Push(&OptCode{
				Type: READIN,
			})
			hasReadin = true
		case '[':
			if last.Type != -1 {
				optCodes.Push(last)
			}
			optCodes.Push(&OptCode{
				Type: BEGINLOOP,
			})
		case ']':
			if last.Type != -1 {
				optCodes.Push(last)
			}
			optCodes.Push(&OptCode{
				Type: ENDLOOP,
			})
		default:
			optCodes.Push(last)
			continue
		}
	}

	processed := OptimizeForLoops(optCodes)

	// Print out go code!
	sb := &strings.Builder{}
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
	for opt := processed.Pop(); processed.HasNext(); opt = processed.Pop() {
		switch opt.Type {
		case PTUP:
			addLine(sb, fmt.Sprintf("ptr += %d", opt.Value), tabs)

		case PTDOWN:
			addLine(sb, fmt.Sprintf("ptr -= %d", opt.Value), tabs)

		case VALUEUP:
			addLine(sb, fmt.Sprintf("buffer[ptr] += byte(%d)", opt.Value), tabs)

		case VALUEDOWN:
			addLine(sb, fmt.Sprintf("buffer[ptr] -= byte(%d)", opt.Value), tabs)

		case READOUT:
			addLine(sb, "fmt.Print(string(buffer[ptr]))", tabs)

		case READIN:
			addLine(sb, "os.Stdin.Read(b)", tabs)
			addLine(sb, "buffer[ptr] = b[0]", tabs)

		case SETZERO:
			addLine(sb, "buffer[ptr] = 0", tabs)

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

	err = ioutil.WriteFile(outputFile, []byte(sb.String()), 0644)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func OptimizeForLoops(optCodes *FIFO) *FIFO {
	output := NewFifo()
	for opt := optCodes.Pop(); optCodes.HasNext(); opt = optCodes.Pop() {
		// This optimization is quite simple.  We are looking for what is essentially [-]
		// This resets the value at current pointer to 0, so we are looking for this
		// and just setting the value to 0
		if opt.Type == BEGINLOOP {
			opt1 := optCodes.Pop()
			opt2 := optCodes.Pop()
			if opt1.Type == VALUEDOWN && opt2.Type == ENDLOOP {
				output.Push(&OptCode{Type: SETZERO})
				continue
			} else {
				// Not a reset, put the values back onto the queue like it never happened
				optCodes.PushHead(opt2)
				optCodes.PushHead(opt1)
			}
		}
		output.Push(opt)
	}
	return output
}

func addLine(sb *strings.Builder, value string, tabs int) {
	for i := 0; i < tabs; i++ {
		sb.WriteString("\t")
	}
	sb.WriteString(value)
	sb.WriteString("\n")
}

type FIFO struct {
	queue []*OptCode
}

func NewFifo() *FIFO {
	return &FIFO{
		queue: make([]*OptCode, 0),
	}
}

func (f *FIFO) Push(v *OptCode) {
	f.queue = append(f.queue, v)
}

func (f *FIFO) PushHead(v *OptCode) {
	f.queue = append([]*OptCode{v}, f.queue...)
}

func (f *FIFO) Pop() *OptCode {
	if len(f.queue) < 1 {
		return nil
	}
	v := f.queue[0]
	f.queue = f.queue[1:]
	return v
}

func (f *FIFO) HasNext() bool {
	return len(f.queue) > 0
}

func (f *FIFO) Peak() *OptCode {
	return f.queue[0]
}

func (f *FIFO) PeakOffset(off int) *OptCode {
	if off > len(f.queue) {
		return nil
	}
	return f.queue[off]
}

func (f *FIFO) PeakLast() *OptCode {
	return f.queue[len(f.queue)-1]
}

func (f *FIFO) PopLast() *OptCode {
	v := f.queue[len(f.queue)-1]
	f.queue = f.queue[:len(f.queue)-1]
	return v
}
