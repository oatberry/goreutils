package main

import "os"

const bufsize = 4096

func main() {
	buffer := make([]byte, bufsize)
	buffer[0] = 'y'
	buffer[1] = '\n'
	for i := 2; i < bufsize; i *= 2 {
		copy(buffer[i:], buffer[:i])
	}

	for {
		os.Stdout.Write(buffer)
	}
}
