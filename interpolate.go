package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"math"
	"strconv"
	"strings"
	"os"
)

// Interpolate reads CSV data from inFile and outputs it to outFile.
// Any missing values in inFile will be replaced with interpolated
// values generated from adjacent data.
// All rows in inFile must have the same number of rows.
func Interpolate(inFile io.Reader, outFile io.Writer) {
	rdr := prepCsvReader(inFile)
	eof := false
	var prevLine []float64
	var curLine []float64
	var nextLine []float64
	lineNum := 0
	terminator := "\n"

	for !eof {
		lineNum++
		fieldCount := len(curLine)
		if fieldCount == 0 {
			fieldCount = 10
		}
		nextLine, eof = readLine(rdr, fieldCount, lineNum)
		outLine := interpolate(prevLine, curLine, nextLine)
		if eof {
			terminator = ""
		}
		outputLine(outFile, outLine, terminator)
		prevLine = curLine
		curLine = nextLine
	}
	outLine := interpolate(curLine, nextLine, nil)
	fmt.Printf("Outputting last line")
	outputLine(outFile, outLine, terminator)
}

func interpolate(prevLine []float64, curLine []float64, nextLine []float64) []float64 {
	if len(curLine) == 0 {
		return nil
	}
	// If lines are long enough for this memory to be an issue it could be skipped for lines without NaNs
	out := make([]float64, len(curLine), len(curLine))
	copy(out[:], curLine)
	for i, v := range out {
		if math.IsNaN(v) {
			var total float64
			var count int
			if i > 0 && !math.IsNaN(curLine[i - 1]) {
				total += curLine[i - 1]
				count++
			}
			if i < len(curLine) - 1 && !math.IsNaN(curLine[i + 1]) {
				total += curLine[i + 1]
				count++
			}
			if len(prevLine) > i && !math.IsNaN(prevLine[i]) {
				total += prevLine[i]
				count++
			}
			if len(nextLine) > i && !math.IsNaN(nextLine[i]) {
				total += nextLine[i]
				count++
			}
			if count > 0 {
				out[i] = total / float64(count)
			}
		}
	}	
	return out
}

func outputLine(outFile io.Writer, line []float64, terminator string)  {
	if line != nil {
		for i, v := range line {
			if i == len(line) - 1 {
				fmt.Fprintf(outFile, "%.9g%s", v, terminator)
			} else {
				fmt.Fprintf(outFile, "%.9g,", v)
			}
		}
	}
}

func prepCsvReader(inFile io.Reader) *csv.Reader {
	rdr := csv.NewReader(inFile)
	rdr.ReuseRecord = true
	return rdr
}

func readLine(rdr *csv.Reader, expectedSize int, lineNum int) ([]float64, bool) {
	result := make([]float64, 0, expectedSize)
	eof := false

	record, err := rdr.Read()
	if err != nil {
		if err == io.EOF {
			eof = true	
		} else if err, ok := err.(*csv.ParseError); ok && err.Err == csv.ErrFieldCount {
			// The input isn't rectangular, some lines have more/less fields than they should
			// Force all lines to have the same number of fields
			if len(record) > expectedSize {
				fmt.Fprintf(os.Stderr, "[%d] Invalid field count: got %d, expected %d, truncating %d field(s)\n",
					lineNum, len(record), expectedSize, len(record) - expectedSize)
				record = record[:expectedSize]
			}
			for len(record) < expectedSize {
				fmt.Fprintf(os.Stderr, "[%d] Invalid field count: got %d, expected %d, appending %d NaN field(s)\n",
					lineNum, len(record), expectedSize, expectedSize - len(record))
				record = append(record, "nan")
			}
		} else {
			fmt.Printf("%#v", err)
			panic(err)
		}
	}

	for i := 0; i < len(record); i++ {
		tok := strings.TrimSpace(record[i])

		result = append(result, 0.0)		
		result[i], err = strconv.ParseFloat(tok, 64)
		if err != nil {
			if "nan" != tok {
				fmt.Fprintf(os.Stderr, "[%d, %d] Invalid text value: \"%s\" (treating it as \"nan\" anyway)\n",
					lineNum,
					i,
					tok)
			}
			result[i] = math.NaN()
		}
	}
	return result, eof
}
