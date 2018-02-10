package main
import "fmt"
func main() {
	buffer := make([]byte, 30000000)
	ptr := 0
	buffer[ptr] = buffer[ptr] + byte(8)
	for buffer[ptr] != 0 {
		ptr = ptr + 1
		buffer[ptr] = buffer[ptr] + byte(4)
		for buffer[ptr] != 0 {
			ptr = ptr + 1
			buffer[ptr] = buffer[ptr] + byte(2)
			ptr = ptr + 1
			buffer[ptr] = buffer[ptr] + byte(3)
			ptr = ptr + 1
			buffer[ptr] = buffer[ptr] + byte(3)
			ptr = ptr + 1
			buffer[ptr] = buffer[ptr] + byte(1)
			ptr = ptr - 4
			buffer[ptr] = buffer[ptr] - byte(1)
		}
		ptr = ptr + 1
		buffer[ptr] = buffer[ptr] + byte(1)
		ptr = ptr + 1
		buffer[ptr] = buffer[ptr] + byte(1)
		ptr = ptr + 1
		buffer[ptr] = buffer[ptr] - byte(1)
		ptr = ptr + 2
		buffer[ptr] = buffer[ptr] + byte(1)
		for buffer[ptr] != 0 {
			ptr = ptr - 1
		}
		ptr = ptr - 1
		buffer[ptr] = buffer[ptr] - byte(1)
	}
	ptr = ptr + 2
	fmt.Print(string(buffer[ptr]))
	ptr = ptr + 1
	buffer[ptr] = buffer[ptr] - byte(3)
	fmt.Print(string(buffer[ptr]))
	buffer[ptr] = buffer[ptr] + byte(7)
	fmt.Print(string(buffer[ptr]))
	fmt.Print(string(buffer[ptr]))
	buffer[ptr] = buffer[ptr] + byte(3)
	fmt.Print(string(buffer[ptr]))
	ptr = ptr + 2
	fmt.Print(string(buffer[ptr]))
	ptr = ptr - 1
	buffer[ptr] = buffer[ptr] - byte(1)
	fmt.Print(string(buffer[ptr]))
	ptr = ptr - 1
	fmt.Print(string(buffer[ptr]))
	buffer[ptr] = buffer[ptr] + byte(3)
	fmt.Print(string(buffer[ptr]))
	buffer[ptr] = buffer[ptr] - byte(6)
	fmt.Print(string(buffer[ptr]))
	buffer[ptr] = buffer[ptr] - byte(8)
	fmt.Print(string(buffer[ptr]))
	ptr = ptr + 2
	buffer[ptr] = buffer[ptr] + byte(1)
	fmt.Print(string(buffer[ptr]))
	ptr = ptr + 1
	buffer[ptr] = buffer[ptr] + byte(2)
	fmt.Print(string(buffer[ptr]))
}
