package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
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
	for _, filename := range filenames {
		file, err := ioutil.ReadFile(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "cat: %v\n", err)
		}

		_, err = os.Stdout.Write(file)
		if err != nil {
			return fmt.Errorf("writing stdout: %v", err)
		}
	}

	return nil
}

func readStdin() error {
	reader := bufio.NewReader(os.Stdin)
	for {
		line, err := reader.ReadSlice(byte('\n'))
		if err == io.EOF {
			return nil
		} else if err != nil {
			return fmt.Errorf("reading stdin: %v", err)
		}

		_, err = os.Stdout.Write(line)
		if err != nil {
			return fmt.Errorf("writing stdout: %v", err)
		}
	}
}
