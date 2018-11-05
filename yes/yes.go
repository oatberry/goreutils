package main

import "os"

const bufsize = 4096

func main() {
	var str []byte
	if len(os.Args) == 1 {
		str = []byte("y\n")
	} else {
		str = []byte(os.Args[1])
		str = append(str, '\n')
	}

	buffer := make([]byte, bufsize)
	copy(buffer[0:len(str)], str)

	var i int
	for i = len(str); i < bufsize; i *= 2 {
		copy(buffer[i:], buffer[:i])
	}

	i /= 2
	for {
		os.Stdout.Write(buffer[:i])
	}
}
