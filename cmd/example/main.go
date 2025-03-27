package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"strings"

	lab2 "github.com/roman-mazur/architecture-lab-2"
)

var (
	expr       = flag.String("e", "", "Expression to compute")
	inputFile  = flag.String("f", "", "File containing the expression")
	outputFile = flag.String("o", "", "File to write the result")
)

func main() {
	flag.Parse()

	if *expr != "" && *inputFile != "" {
		fmt.Fprintln(os.Stderr, "Error: Cannot use both -e and -f options")
		os.Exit(1)
	}

	var inputReader io.Reader
	var err error
	if *inputFile != "" {
		inputReader, err = os.Open(*inputFile)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error opening file: %v\n", err)
			os.Exit(1)
		}
		defer inputReader.(*os.File).Close()
	} else {
		inputReader = os.Stdin
		if *expr != "" {
			inputReader = strings.NewReader(*expr)
		}
	}
	var outputWriter *os.File
	if *outputFile != "" {
		outputWriter, err = os.Create(*outputFile)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error creating output file: %v\n", err)
			os.Exit(1)
		}
		defer outputWriter.Close()
	} else {
		outputWriter = os.Stdout
	}
	handler := &lab2.ComputeHandler{
		Input:  inputReader,
		Output: outputWriter,
	}
	if err := handler.Compute(); err != nil {
		fmt.Fprintf(os.Stderr, "Computation error: %v\n", err)
		os.Exit(1)
	}
}
