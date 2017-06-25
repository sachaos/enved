package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

func main() {
	var infile, outfile string
	var err interface{}
	flag.StringVar(&outfile, "o", "", "output filename (default STDOUT)")
	flag.Parse()

	infile = flag.Arg(0)

	var reader io.Reader
	if infile != "" {
		reader, err = os.Open(infile)
		if err != nil {
			panic(err)
		}
	} else {
		reader = os.Stdin
	}

	var writer io.Writer
	if outfile != "" {
		writer, err = os.Create(outfile)
		if err != nil {
			panic(err)
		}
	} else {
		writer = os.Stdout
	}

	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		fmt.Fprintln(writer, os.ExpandEnv(scanner.Text()))
	}
}
