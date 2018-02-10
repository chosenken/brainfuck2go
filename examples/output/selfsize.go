package main
import "fmt"
func main() {
	buffer := make([]byte, 30000)
	ptr := 0
	buffer[ptr] = buffer[ptr] + byte(10)
	for buffer[ptr] != 0 {
		ptr = ptr + 1
		buffer[ptr] = buffer[ptr] + byte(5)
		ptr = ptr + 1
		buffer[ptr] = buffer[ptr] + byte(1)
		ptr = ptr + -2
		buffer[ptr] = buffer[ptr] - byte(1)
	}
	ptr = ptr + 1
	buffer[ptr] = buffer[ptr] + byte(1)
	fmt.Print(string(buffer[ptr]))
	buffer[ptr] = buffer[ptr] - byte(2)
	fmt.Print(string(buffer[ptr]))
	ptr = ptr + 1
	fmt.Print(string(buffer[ptr]))
}
