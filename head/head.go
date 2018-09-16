package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

type config struct {
	nLines  int      // number of first lines to print
	nBytes  int      // number of first bytes to print
	verbose bool     // always print headers with file names
	quiet   bool     // never print file headers
	headers bool     // whether to print headers
	files   []string // files to print
}

func main() {
	config := parseArgs()

	err := head(config)
	if err != nil {
		fmt.Fprintf(os.Stderr, "head: %v\n", err)
		os.Exit(1)
	}
}

func parseArgs() config {
	config := config{}

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s [OPTION]... [FILE]...\n", os.Args[0])
		flag.PrintDefaults()
	}

	flag.IntVar(&config.nLines, "n", 10, "number of lines to print")
	flag.IntVar(&config.nBytes, "c", 0, "number of bytes to print")
	flag.BoolVar(&config.verbose, "v", false, "print filename headers")
	flag.BoolVar(&config.quiet, "q", false, "never print filename headers")

	flag.Parse()
	config.files = flag.Args()
	config.headers = !config.quiet && (len(config.files) > 1 || config.verbose)
	return config
}

func head(config config) error {
	bufStdout := bufio.NewWriter(os.Stdout)
	defer bufStdout.Flush()

	for _, filename := range config.files {
		bufStdout.Flush()

		if config.headers {
			fmt.Printf("==> %v <==\n", filename)
		}

		file, err := os.Open(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "head: %v\n", err)
			continue
		}

		if config.nBytes == 0 {
			err = printLines(file, config.nLines, bufStdout)
		} else {
			err = printBytes(file, config.nBytes, bufStdout)
		}

		if config.headers {
			bufStdout.WriteByte('\n')
		}

		if err != nil {
			return err
		}
	}

	return nil
}

func printLines(file *os.File, nLines int, output *bufio.Writer) error {
	bufFile := bufio.NewReader(file)
	var bytes []byte
	var err error
	for i := 0; i < nLines; i++ {
		bytes, err = bufFile.ReadBytes('\n')
		if err == io.EOF {
			return nil
		} else if err != nil {
			return err
		}

		_, err = output.Write(bytes)
		if err != nil {
			return err
		}
	}
	return nil
}

func printBytes(file *os.File, nBytes int, output *bufio.Writer) error {
	bytes := make([]byte, nBytes)
	_, err := file.Read(bytes)
	if err != nil && err != io.EOF {
		return err
	}

	output.Write(bytes)
	return nil
}
