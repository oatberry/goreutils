package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type config struct {
	separator string // what to separate numbers with
	first     uint64
	increment uint64
	last      uint64
}

func main() {
	config := parseArgs()

	err := seq(config.first, config.increment, config.last, config.separator)
	if err != nil {
		fmt.Fprintf(os.Stderr, "seq: %v\n", err)
		os.Exit(1)
	}
}

func seq(first, inc, last uint64, separator string) error {
	bufout := bufio.NewWriter(os.Stdout)
	defer bufout.Flush()

	for i := first; i <= last; i += inc {
		_, err := bufout.Write([]byte(fmt.Sprint(i, separator)))
		if err != nil {
			return err
		}
	}

	if !strings.HasSuffix(separator, "\n") {
		bufout.WriteByte('\n')
	}

	return nil
}

func parseArgs() config {
	config := config{}
	config.separator = "\n"
	config.first = 1
	config.increment = 1

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s [OPTION]... [FIRST] [INCREMENT] <LAST>\n", os.Args[0])
		flag.PrintDefaults()
	}

	flag.StringVar(&config.separator, "s", "\n", "use STRING to separate numbers")
	flag.Parse()

	var err error
	endArgs := flag.Args()
	switch len(endArgs) {
	case 0:
		fmt.Fprintf(os.Stderr, "seq: missing argument\n")
		os.Exit(1)
	case 1:
		config.last, err = strconv.ParseUint(endArgs[0], 10, 64)
	case 2:
		config.first, err = strconv.ParseUint(endArgs[0], 10, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "seq: %v\n", err)
			os.Exit(1)
		}
		config.last, err = strconv.ParseUint(endArgs[1], 10, 64)
	case 3:
		config.first, err = strconv.ParseUint(endArgs[0], 10, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "seq: %v\n", err)
			os.Exit(1)
		}
		config.increment, err = strconv.ParseUint(endArgs[1], 10, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "seq: %v\n", err)
			os.Exit(1)
		}
		config.last, err = strconv.ParseUint(endArgs[2], 10, 64)
	}

	if err != nil {
		fmt.Fprintf(os.Stderr, "seq: %v\n", err)
		os.Exit(1)
	}

	return config
}
