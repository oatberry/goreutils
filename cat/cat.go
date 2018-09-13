package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	var err error
	if len(os.Args) > 1 {
		err = catFiles(os.Args[1:])
	} else {
		err = readStdin()
	}

	if err != nil {
		fmt.Fprintf(os.Stderr, "cat: error %v\n", err)
		os.Exit(1)
	}
}

func catFiles(filenames []string) error {
	bufStdout := bufio.NewWriterSize(os.Stdout, 128000) // GNU cat uses 128KB buffer
	for _, filename := range filenames {
		file, err := os.Open(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "cat: %v\n", err)
		}
		defer file.Close()

		bufStdout.ReadFrom(file)
	}

	bufStdout.Flush()
	return nil
}

func readStdin() error {
	bufStdin := bufio.NewReader(os.Stdin)
	for {
		line, err := bufStdin.ReadSlice(byte('\n'))
		if err == io.EOF {
			return nil
		} else if err != nil {
			return fmt.Errorf("reading stdin: %v", err)
		}

		os.Stdout.Write(line)
	}
}
