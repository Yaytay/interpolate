package main

import (
	"flag"
	"fmt"
	"io"
	"os"
)

func displayUsage() {
    fmt.Fprintf(os.Stderr, "%s usage: \n", os.Args[0])
    fmt.Fprintf(os.Stderr, "%s [OPTIONS] <inputfile> <outputfile>\n", os.Args[0])
    fmt.Fprintf(os.Stderr, " inputfile may be \"-\" to input from stdin\n")
    fmt.Fprintf(os.Stderr, " outputfile may be \"-\" to output to stderr\n")
    fmt.Fprintf(os.Stderr, " options:\n")
    flag.PrintDefaults()
    os.Exit(1)
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	help := flag.Bool("?", false, "Print this message")

	flag.Parse()

	if *help || len(flag.Args()) != 2 {
		displayUsage()
	}


	inFilename := flag.Args()[0]
	var inFile io.Reader
	if "-" == inFilename {
		inFile = os.Stdin
	} else {
		f, err := os.Open(inFilename)
		check(err)
		defer f.Close()
		inFile = f
	}

	outFilename := flag.Args()[1]
	var outFile io.Writer
	if "-" == outFilename {
		outFile = os.Stdout
	} else {
		f, err := os.Create(outFilename)
		check(err)
		defer f.Close()
		outFile = f
	}

	Interpolate(inFile, outFile)

}
