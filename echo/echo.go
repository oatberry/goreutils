package main

import (
	"bufio"
	"os"
)

func main() {
	bufStdout := bufio.NewWriter(os.Stdout)
	argslen := len(os.Args) - 2
	for i, arg := range os.Args[1:] {
		bufStdout.WriteString(arg)
		// avoid trailing space
		if i < argslen {
			bufStdout.WriteRune(' ')
		}
	}

	bufStdout.WriteRune('\n')
	bufStdout.Flush()
}
